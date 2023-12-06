const { test, expect } = require("@playwright/test");
const { start, stop, restart, cleanRestart } = require("./evcc");

const CONFIG_EMPTY = "config-empty.evcc.yaml";
const CONFIG_WITH_VEHICLE = "config-with-vehicle.evcc.yaml";

test.beforeAll(async () => {
  await start(CONFIG_EMPTY);
});
test.afterAll(async () => {
  await stop();
});

test.describe("basics", async () => {
  test("navigation to config", async ({ page }) => {
    await page.goto("/");
    await page.getByTestId("topnavigation-button").click();
    await page.getByRole("button", { name: "Settings" }).click();
    await page.getByLabel("Experimental 🧪").click();
    await page.getByRole("button", { name: "Close" }).click();
    await page.getByTestId("topnavigation-button").click();
    await page.getByRole("link", { name: "Configuration" }).click();
    await expect(page.getByRole("heading", { name: "Configuration" })).toBeVisible();
  });
  test("alert box should always be visible", async ({ page }) => {
    await page.goto("/#/config");
    await expect(page.getByRole("alert")).toBeVisible();
  });
});

test.describe("vehicles", async () => {
  test("create, edit and delete vehicles", async ({ page }) => {
    await page.goto("/#/config");

    await expect(page.getByTestId("vehicle")).toHaveCount(0);
    const vehicleModal = page.getByTestId("vehicle-modal");

    // create #1
    await page.getByTestId("add-vehicle").click();
    await vehicleModal.getByLabel("Manufacturer").selectOption("Generic vehicle");
    await vehicleModal.getByLabel("Title").fill("Green Car");
    await vehicleModal.getByRole("button", { name: "Validate & save" }).click();

    await expect(page.getByTestId("vehicle")).toHaveCount(1);

    // create #2
    await page.getByTestId("add-vehicle").click();
    await vehicleModal.getByLabel("Manufacturer").selectOption("Generic vehicle");
    await vehicleModal.getByLabel("Title").fill("Yellow Van");
    await vehicleModal.getByRole("button", { name: "Validate & save" }).click();

    await expect(page.getByTestId("vehicle")).toHaveCount(2);
    await expect(page.getByTestId("vehicle").nth(0)).toHaveText(/Green Car/);
    await expect(page.getByTestId("vehicle").nth(1)).toHaveText(/Yellow Van/);

    // edit #1
    await page.getByTestId("vehicle").nth(0).getByRole("button", { name: "edit" }).click();
    await expect(vehicleModal.getByLabel("Title")).toHaveValue("Green Car");
    await vehicleModal.getByLabel("Title").fill("Fancy Car");
    await vehicleModal.getByRole("button", { name: "Validate & save" }).click();

    await expect(page.getByTestId("vehicle")).toHaveCount(2);
    await expect(page.getByTestId("vehicle").nth(0)).toHaveText(/Fancy Car/);

    // delete #1
    await page.getByTestId("vehicle").nth(0).getByRole("button", { name: "edit" }).click();
    await vehicleModal.getByRole("button", { name: "Delete Vehicle" }).click();

    await expect(page.getByTestId("vehicle")).toHaveCount(1);
    await expect(page.getByTestId("vehicle").nth(0)).toHaveText(/Yellow Van/);

    // delete #2
    await page.getByTestId("vehicle").nth(0).getByRole("button", { name: "edit" }).click();
    await vehicleModal.getByRole("button", { name: "Delete Vehicle" }).click();

    await expect(page.getByTestId("vehicle")).toHaveCount(0);
  });

  test("config should survive restart", async ({ page }) => {
    await page.goto("/#/config");

    await expect(page.getByTestId("vehicle")).toHaveCount(0);
    const vehicleModal = page.getByTestId("vehicle-modal");

    // create #1 & #2
    await page.getByTestId("add-vehicle").click();
    await vehicleModal.getByLabel("Manufacturer").selectOption("Generic vehicle");
    await vehicleModal.getByLabel("Title").fill("Green Car");
    await vehicleModal.getByRole("button", { name: "Validate & save" }).click();

    await page.getByTestId("add-vehicle").click();
    await vehicleModal.getByLabel("Manufacturer").selectOption("Generic vehicle");
    await vehicleModal.getByLabel("Title").fill("Yellow Van");
    await vehicleModal.getByLabel("car").click();
    await vehicleModal.getByLabel("van").check();
    await vehicleModal.getByRole("button", { name: "Validate & save" }).click();

    await expect(page.getByTestId("vehicle")).toHaveCount(2);

    // restart evcc
    await restart(CONFIG_EMPTY);
    await page.reload();

    await expect(page.getByTestId("vehicle")).toHaveCount(2);
    await expect(page.getByTestId("vehicle").nth(0)).toHaveText(/Green Car/);
    await expect(page.getByTestId("vehicle").nth(1)).toHaveText(/Yellow Van/);
  });

  test("mixed config (yaml + db)", async ({ page }) => {
    await cleanRestart(CONFIG_WITH_VEHICLE);

    await page.goto("/#/config");

    await expect(page.getByTestId("vehicle")).toHaveCount(1);
    const vehicleModal = page.getByTestId("vehicle-modal");

    // create #1
    await page.getByTestId("add-vehicle").click();
    await vehicleModal.getByLabel("Manufacturer").selectOption("Generic vehicle");
    await vehicleModal.getByLabel("Title").fill("Green Car");
    await vehicleModal.getByRole("button", { name: "Validate & save" }).click();

    await expect(page.getByTestId("vehicle")).toHaveCount(2);
    await expect(page.getByTestId("vehicle").nth(0)).toHaveText(/YAML Bike/);
    await expect(page.getByTestId("vehicle").nth(1)).toHaveText(/Green Car/);
  });
});
