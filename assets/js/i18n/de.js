export default {
  header: {
    docs: "Dokumentation",
    blog: "Blog",
    github: "GitHub",
    login: "Fahrzeug Logins",
    about: "Über evcc",
  },
  footer: {
    version: {
      availableShort: "Update",
      availableLong: "Update verfügbar",
      modalTitle: "Update verfügbar",
      modalUpdateStarted: "Nach der Aktualisierung wird evcc neu gestartet.",
      modalInstalledVersion: "Aktuell installierte Version",
      modalNoReleaseNotes:
        "Keine Releasenotes verfügbar. Mehr Informationen zur neuen Version findest du hier:",
      modalCancel: "Abbrechen",
      modalUpdate: "Aktualisieren",
      modalUpdateNow: "Jetzt aktualisieren",
      modalDownload: "Download",
      modalUpdateStatusStart: "Aktualisierung gestartet: ",
      modalUpdateStatusFailed: "Aktualisierung nicht möglich: ",
    },
    savings: {
      footerShort: "{percent}% Sonne",
      footerLong: "{percent}% Sonnenenergie",
      modalTitle: "Auswertung Ladeenergie",
      sinceServerStart: "Seit Serverstart {since}.",
      percentTitle: "Sonnenenergie",
      percentSelf: "{self} kWh Sonne",
      percentGrid: "{grid} kWh Netz",
      priceTitle: "Energiepreis",
      priceFeedIn: "{feedInPrice} Einpeisung",
      priceGrid: "{gridPrice} Netz",
      savingsTitle: "Ersparnis",
      savingsComparedToGrid: "gegenüber Netzbezug",
      savingsTotalEnergy: "{total} kWh geladen",
    },
    sponsor: {
      thanks: "Danke für deine Unterstützung, {sponsor}! Das hilft uns bei der Weiterentwicklung.",
      confetti: "Lust auf Konfetti?",
      supportUs:
        "Unsere Mission: Sonne tanken zum Standard machen. Helfe uns und unterstütze evcc finanziell.",
      sticker: "...oder evcc Sticker?",
      confettiPromise: "Es gibt auch Sticker und digitales Konfetti ;)",
      becomeSponsor: "Sponsor werden",
    },
  },
  notifications: {
    modalTitle: "Meldungen",
    dismissAll: "Meldungen entfernen",
  },
  main: {
    energyflow: {
      noEnergy: "Kein Energiefluss",
      homePower: "Verbrauch",
      pvProduction: "Erzeugung",
      loadpoints: "Ladepunkt | Ladepunkt | {count} Ladepunkte",
      battery: "Batterie",
      batteryCharge: "Batterie laden",
      batteryDischarge: "Batterie entladen",
      gridImport: "Netzbezug",
      selfConsumption: "Eigenverbrauch",
      pvExport: "Einspeisung",
    },
    mode: {
      off: "Aus",
      minpv: "Min+PV",
      pv: "PV",
      now: "Schnell",
    },
    loadpoint: {
      fallbackName: "Ladepunkt",
      remoteDisabledSoft: "{source}: Adaptives PV-Laden deaktiviert",
      remoteDisabledHard: "{source}: Deaktiviert",
      power: "Leistung",
      charged: "Geladen",
      duration: "Dauer",
      remaining: "Restzeit",
    },
    vehicles: "Parkplatz",
    vehicle: {
      fallbackName: "Fahrzeug",
      vehicleSoC: "Ladestand",
      targetSoC: "Ladeziel",
      none: "Kein Fahrzeug",
      unknown: "Gastfahrzeug",
    },
    vehicleSoC: {
      disconnected: "getrennt",
      charging: "lädt",
      ready: "bereit",
      connected: "verbunden",
    },
    vehicleStatus: {
      minCharge: "Mindestladung bis {soc}%.",
      waitForVehicle: "Ladebereit. Warte auf Fahrzeug.",
      charging: "Ladevorgang aktiv.",
      targetChargePlanned: "Zielladen geplant. Ladung startet {time} Uhr.",
      targetChargeWaitForVehicle: "Zielladen bereit. Warte auf Fahrzeug.",
      targetChargeActive: "Zielladen aktiv.",
      connected: "Verbunden.",
      pvDisable: "Zu wenig Überschuss. Pausiere in {remaining}.",
      pvEnable: "Überschuss verfügbar. Starte in {remaining}.",
      scale1p: "Reduziere auf einphasig in {remaining}.",
      scale3p: "Erhöhe auf dreiphasig in {remaining}.",
      disconnected: "Nicht verbunden.",
      unknown: "",
    },
    provider: {
      login: "anmelden",
      logout: "abmelden",
    },
    targetCharge: {
      title: "Zielzeit",
      inactiveLabel: "Zielzeit",
      activeLabel: "{time}",
      modalTitle: "Zielzeit festlegen",
      setTargetTime: "keine",
      description: "Wann soll das Fahrzeug auf {targetSoC}% geladen sein?",
      today: "heute",
      tomorrow: "morgen",
      targetIsInThePast: "Zeitpunkt liegt in der Vergangenheit.",
      remove: "Entfernen",
      activate: "Aktivieren",
      experimentalLabel: "Experimentell",
      experimentalText: `
        Dieses Feature funktioniert, ist aber noch nicht perfekt.
        Bitte melde unerwartetes Verhalten in unseren
      `,
    },
  },
  offline: {
    message: "Keine Verbindung zum Server.",
    reload: "Reload?",
  },
};
