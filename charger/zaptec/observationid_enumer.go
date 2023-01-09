// Code generated by "enumer -type ObservationID"; DO NOT EDIT.

package zaptec

import (
	"fmt"
	"strings"
)

const _ObservationIDName = "IsOcppConnectedIsOnlinePulseUnknownOfflineModeAuthenticationRequiredPaymentActivePaymentCurrencyPaymentSessionUnitPricePaymentEnergyUnitPricePaymentTimeUnitPriceCommunicationModePermanentCableLockProductCodeHmiBrightnessLockCableWhenConnectedSoftStartDisabledFirmwareApiHostMIDBlinkEnabledTemperatureInternal5TemperatureInternal6TemperatureInternalLimitTemperatureInternalMaxLimitHumidityVoltagePhase1VoltagePhase2VoltagePhase3CurrentPhase1CurrentPhase2CurrentPhase3ChargerMaxCurrentChargerMinCurrentActivePhasesTotalChargePowerRcdCurrentInternal12vCurrentPowerFactorSetPhasesMaxPhasesChargerOfflinePhaseChargerOfflineCurrentRcdCalibrationRcdCalibrationNoiseTotalChargePowerSessionSignedMeterValueSignedMeterValueIntervalSessionEnergyCountExportActiveSessionEnergyCountExportReactiveSessionEnergyCountImportActiveSessionEnergyCountImportReactiveSoftStartTimeChargeDurationChargeModeChargePilotLevelInstantChargePilotLevelAveragePilotVsProximityTimeChargeCurrentInstallationMaxLimitChargeCurrentSetChargerOperationModeIsEnabledIsStandAloneChargerCurrentUserUuidDeprecatedCableTypeNetworkTypeDetectedCarGridTestResultFinalStopActiveSessionIdentifierChargerCurrentUserUuidCompletedSessionNewChargeCardAuthenticationListVersionEnabledNfcTechnologiesLteRoamingDisabledInstallationIdRoutingIdNotificationsWarningsDiagnosticsModeInternalDiagnosticsLogDiagnosticsStringCommunicationSignalStrengthCloudConnectionStatusMcuResetSourceMcuRxErrorsMcuToVariscitePacketErrorsVarisciteToMcuPacketErrorsUptimeVarisciteUptimeMCUCarSessionLogCommunicationModeConfigurationInconsistencyRawPilotMonitorIT3PhaseDiagnosticsLogPilotTestResultsUnconditionalNfcDetectionIndicationEmcTestCounterProductionTestResultsPostProductionTestResultsSmartMainboardSoftwareApplicationVersionSmartMainboardSoftwareBootloaderVersionSmartComputerSoftwareApplicationVersionSmartComputerSoftwareBootloaderVersionSmartComputerHardwareVersionMacMainMacPlcModuleGridMacWiFiMacPlcModuleEvLteImsiLteMsisdnLteIccidLteImeiMIDCalibration"
const _ObservationIDLowerName = "isocppconnectedisonlinepulseunknownofflinemodeauthenticationrequiredpaymentactivepaymentcurrencypaymentsessionunitpricepaymentenergyunitpricepaymenttimeunitpricecommunicationmodepermanentcablelockproductcodehmibrightnesslockcablewhenconnectedsoftstartdisabledfirmwareapihostmidblinkenabledtemperatureinternal5temperatureinternal6temperatureinternallimittemperatureinternalmaxlimithumidityvoltagephase1voltagephase2voltagephase3currentphase1currentphase2currentphase3chargermaxcurrentchargermincurrentactivephasestotalchargepowerrcdcurrentinternal12vcurrentpowerfactorsetphasesmaxphaseschargerofflinephasechargerofflinecurrentrcdcalibrationrcdcalibrationnoisetotalchargepowersessionsignedmetervaluesignedmetervalueintervalsessionenergycountexportactivesessionenergycountexportreactivesessionenergycountimportactivesessionenergycountimportreactivesoftstarttimechargedurationchargemodechargepilotlevelinstantchargepilotlevelaveragepilotvsproximitytimechargecurrentinstallationmaxlimitchargecurrentsetchargeroperationmodeisenabledisstandalonechargercurrentuseruuiddeprecatedcabletypenetworktypedetectedcargridtestresultfinalstopactivesessionidentifierchargercurrentuseruuidcompletedsessionnewchargecardauthenticationlistversionenablednfctechnologieslteroamingdisabledinstallationidroutingidnotificationswarningsdiagnosticsmodeinternaldiagnosticslogdiagnosticsstringcommunicationsignalstrengthcloudconnectionstatusmcuresetsourcemcurxerrorsmcutovariscitepacketerrorsvariscitetomcupacketerrorsuptimevarisciteuptimemcucarsessionlogcommunicationmodeconfigurationinconsistencyrawpilotmonitorit3phasediagnosticslogpilottestresultsunconditionalnfcdetectionindicationemctestcounterproductiontestresultspostproductiontestresultssmartmainboardsoftwareapplicationversionsmartmainboardsoftwarebootloaderversionsmartcomputersoftwareapplicationversionsmartcomputersoftwarebootloaderversionsmartcomputerhardwareversionmacmainmacplcmodulegridmacwifimacplcmoduleevlteimsiltemsisdnlteiccidlteimeimidcalibration"

var _ObservationIDMap = map[ObservationID]string{
	-3:  _ObservationIDName[0:15],
	-2:  _ObservationIDName[15:23],
	-1:  _ObservationIDName[23:28],
	0:   _ObservationIDName[28:35],
	1:   _ObservationIDName[35:46],
	120: _ObservationIDName[46:68],
	130: _ObservationIDName[68:81],
	131: _ObservationIDName[81:96],
	132: _ObservationIDName[96:119],
	133: _ObservationIDName[119:141],
	134: _ObservationIDName[141:161],
	150: _ObservationIDName[161:178],
	151: _ObservationIDName[178:196],
	152: _ObservationIDName[196:207],
	153: _ObservationIDName[207:220],
	154: _ObservationIDName[220:242],
	155: _ObservationIDName[242:259],
	156: _ObservationIDName[259:274],
	170: _ObservationIDName[274:289],
	201: _ObservationIDName[289:309],
	202: _ObservationIDName[309:329],
	203: _ObservationIDName[329:353],
	241: _ObservationIDName[353:380],
	270: _ObservationIDName[380:388],
	501: _ObservationIDName[388:401],
	502: _ObservationIDName[401:414],
	503: _ObservationIDName[414:427],
	507: _ObservationIDName[427:440],
	508: _ObservationIDName[440:453],
	509: _ObservationIDName[453:466],
	510: _ObservationIDName[466:483],
	511: _ObservationIDName[483:500],
	512: _ObservationIDName[500:512],
	513: _ObservationIDName[512:528],
	515: _ObservationIDName[528:538],
	517: _ObservationIDName[538:556],
	518: _ObservationIDName[556:567],
	519: _ObservationIDName[567:576],
	520: _ObservationIDName[576:585],
	522: _ObservationIDName[585:604],
	523: _ObservationIDName[604:625],
	540: _ObservationIDName[625:639],
	541: _ObservationIDName[639:658],
	553: _ObservationIDName[658:681],
	554: _ObservationIDName[681:697],
	555: _ObservationIDName[697:721],
	560: _ObservationIDName[721:751],
	561: _ObservationIDName[751:783],
	562: _ObservationIDName[783:813],
	563: _ObservationIDName[813:845],
	570: _ObservationIDName[845:858],
	701: _ObservationIDName[858:872],
	702: _ObservationIDName[872:882],
	703: _ObservationIDName[882:905],
	704: _ObservationIDName[905:928],
	706: _ObservationIDName[928:948],
	707: _ObservationIDName[948:981],
	708: _ObservationIDName[981:997],
	710: _ObservationIDName[997:1017],
	711: _ObservationIDName[1017:1026],
	712: _ObservationIDName[1026:1038],
	713: _ObservationIDName[1038:1070],
	714: _ObservationIDName[1070:1079],
	715: _ObservationIDName[1079:1090],
	716: _ObservationIDName[1090:1101],
	717: _ObservationIDName[1101:1115],
	718: _ObservationIDName[1115:1130],
	721: _ObservationIDName[1130:1147],
	722: _ObservationIDName[1147:1169],
	723: _ObservationIDName[1169:1185],
	750: _ObservationIDName[1185:1198],
	751: _ObservationIDName[1198:1223],
	752: _ObservationIDName[1223:1245],
	753: _ObservationIDName[1245:1263],
	800: _ObservationIDName[1263:1277],
	801: _ObservationIDName[1277:1286],
	803: _ObservationIDName[1286:1299],
	804: _ObservationIDName[1299:1307],
	805: _ObservationIDName[1307:1322],
	807: _ObservationIDName[1322:1344],
	808: _ObservationIDName[1344:1361],
	809: _ObservationIDName[1361:1388],
	810: _ObservationIDName[1388:1409],
	811: _ObservationIDName[1409:1423],
	812: _ObservationIDName[1423:1434],
	813: _ObservationIDName[1434:1460],
	814: _ObservationIDName[1460:1486],
	820: _ObservationIDName[1486:1501],
	821: _ObservationIDName[1501:1510],
	850: _ObservationIDName[1510:1523],
	851: _ObservationIDName[1523:1566],
	852: _ObservationIDName[1566:1581],
	853: _ObservationIDName[1581:1603],
	854: _ObservationIDName[1603:1619],
	855: _ObservationIDName[1619:1654],
	899: _ObservationIDName[1654:1668],
	900: _ObservationIDName[1668:1689],
	901: _ObservationIDName[1689:1714],
	908: _ObservationIDName[1714:1754],
	909: _ObservationIDName[1754:1793],
	911: _ObservationIDName[1793:1832],
	912: _ObservationIDName[1832:1870],
	913: _ObservationIDName[1870:1898],
	950: _ObservationIDName[1898:1905],
	951: _ObservationIDName[1905:1921],
	952: _ObservationIDName[1921:1928],
	953: _ObservationIDName[1928:1942],
	960: _ObservationIDName[1942:1949],
	961: _ObservationIDName[1949:1958],
	962: _ObservationIDName[1958:1966],
	963: _ObservationIDName[1966:1973],
	980: _ObservationIDName[1973:1987],
}

func (i ObservationID) String() string {
	if str, ok := _ObservationIDMap[i]; ok {
		return str
	}
	return fmt.Sprintf("ObservationID(%d)", i)
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ObservationIDNoOp() {
	var x [1]struct{}
	_ = x[IsOcppConnected-(-3)]
	_ = x[IsOnline-(-2)]
	_ = x[Pulse-(-1)]
	_ = x[Unknown-(0)]
	_ = x[OfflineMode-(1)]
	_ = x[AuthenticationRequired-(120)]
	_ = x[PaymentActive-(130)]
	_ = x[PaymentCurrency-(131)]
	_ = x[PaymentSessionUnitPrice-(132)]
	_ = x[PaymentEnergyUnitPrice-(133)]
	_ = x[PaymentTimeUnitPrice-(134)]
	_ = x[CommunicationMode-(150)]
	_ = x[PermanentCableLock-(151)]
	_ = x[ProductCode-(152)]
	_ = x[HmiBrightness-(153)]
	_ = x[LockCableWhenConnected-(154)]
	_ = x[SoftStartDisabled-(155)]
	_ = x[FirmwareApiHost-(156)]
	_ = x[MIDBlinkEnabled-(170)]
	_ = x[TemperatureInternal5-(201)]
	_ = x[TemperatureInternal6-(202)]
	_ = x[TemperatureInternalLimit-(203)]
	_ = x[TemperatureInternalMaxLimit-(241)]
	_ = x[Humidity-(270)]
	_ = x[VoltagePhase1-(501)]
	_ = x[VoltagePhase2-(502)]
	_ = x[VoltagePhase3-(503)]
	_ = x[CurrentPhase1-(507)]
	_ = x[CurrentPhase2-(508)]
	_ = x[CurrentPhase3-(509)]
	_ = x[ChargerMaxCurrent-(510)]
	_ = x[ChargerMinCurrent-(511)]
	_ = x[ActivePhases-(512)]
	_ = x[TotalChargePower-(513)]
	_ = x[RcdCurrent-(515)]
	_ = x[Internal12vCurrent-(517)]
	_ = x[PowerFactor-(518)]
	_ = x[SetPhases-(519)]
	_ = x[MaxPhases-(520)]
	_ = x[ChargerOfflinePhase-(522)]
	_ = x[ChargerOfflineCurrent-(523)]
	_ = x[RcdCalibration-(540)]
	_ = x[RcdCalibrationNoise-(541)]
	_ = x[TotalChargePowerSession-(553)]
	_ = x[SignedMeterValue-(554)]
	_ = x[SignedMeterValueInterval-(555)]
	_ = x[SessionEnergyCountExportActive-(560)]
	_ = x[SessionEnergyCountExportReactive-(561)]
	_ = x[SessionEnergyCountImportActive-(562)]
	_ = x[SessionEnergyCountImportReactive-(563)]
	_ = x[SoftStartTime-(570)]
	_ = x[ChargeDuration-(701)]
	_ = x[ChargeMode-(702)]
	_ = x[ChargePilotLevelInstant-(703)]
	_ = x[ChargePilotLevelAverage-(704)]
	_ = x[PilotVsProximityTime-(706)]
	_ = x[ChargeCurrentInstallationMaxLimit-(707)]
	_ = x[ChargeCurrentSet-(708)]
	_ = x[ChargerOperationMode-(710)]
	_ = x[IsEnabled-(711)]
	_ = x[IsStandAlone-(712)]
	_ = x[ChargerCurrentUserUuidDeprecated-(713)]
	_ = x[CableType-(714)]
	_ = x[NetworkType-(715)]
	_ = x[DetectedCar-(716)]
	_ = x[GridTestResult-(717)]
	_ = x[FinalStopActive-(718)]
	_ = x[SessionIdentifier-(721)]
	_ = x[ChargerCurrentUserUuid-(722)]
	_ = x[CompletedSession-(723)]
	_ = x[NewChargeCard-(750)]
	_ = x[AuthenticationListVersion-(751)]
	_ = x[EnabledNfcTechnologies-(752)]
	_ = x[LteRoamingDisabled-(753)]
	_ = x[InstallationId-(800)]
	_ = x[RoutingId-(801)]
	_ = x[Notifications-(803)]
	_ = x[Warnings-(804)]
	_ = x[DiagnosticsMode-(805)]
	_ = x[InternalDiagnosticsLog-(807)]
	_ = x[DiagnosticsString-(808)]
	_ = x[CommunicationSignalStrength-(809)]
	_ = x[CloudConnectionStatus-(810)]
	_ = x[McuResetSource-(811)]
	_ = x[McuRxErrors-(812)]
	_ = x[McuToVariscitePacketErrors-(813)]
	_ = x[VarisciteToMcuPacketErrors-(814)]
	_ = x[UptimeVariscite-(820)]
	_ = x[UptimeMCU-(821)]
	_ = x[CarSessionLog-(850)]
	_ = x[CommunicationModeConfigurationInconsistency-(851)]
	_ = x[RawPilotMonitor-(852)]
	_ = x[IT3PhaseDiagnosticsLog-(853)]
	_ = x[PilotTestResults-(854)]
	_ = x[UnconditionalNfcDetectionIndication-(855)]
	_ = x[EmcTestCounter-(899)]
	_ = x[ProductionTestResults-(900)]
	_ = x[PostProductionTestResults-(901)]
	_ = x[SmartMainboardSoftwareApplicationVersion-(908)]
	_ = x[SmartMainboardSoftwareBootloaderVersion-(909)]
	_ = x[SmartComputerSoftwareApplicationVersion-(911)]
	_ = x[SmartComputerSoftwareBootloaderVersion-(912)]
	_ = x[SmartComputerHardwareVersion-(913)]
	_ = x[MacMain-(950)]
	_ = x[MacPlcModuleGrid-(951)]
	_ = x[MacWiFi-(952)]
	_ = x[MacPlcModuleEv-(953)]
	_ = x[LteImsi-(960)]
	_ = x[LteMsisdn-(961)]
	_ = x[LteIccid-(962)]
	_ = x[LteImei-(963)]
	_ = x[MIDCalibration-(980)]
}

var _ObservationIDValues = []ObservationID{IsOcppConnected, IsOnline, Pulse, Unknown, OfflineMode, AuthenticationRequired, PaymentActive, PaymentCurrency, PaymentSessionUnitPrice, PaymentEnergyUnitPrice, PaymentTimeUnitPrice, CommunicationMode, PermanentCableLock, ProductCode, HmiBrightness, LockCableWhenConnected, SoftStartDisabled, FirmwareApiHost, MIDBlinkEnabled, TemperatureInternal5, TemperatureInternal6, TemperatureInternalLimit, TemperatureInternalMaxLimit, Humidity, VoltagePhase1, VoltagePhase2, VoltagePhase3, CurrentPhase1, CurrentPhase2, CurrentPhase3, ChargerMaxCurrent, ChargerMinCurrent, ActivePhases, TotalChargePower, RcdCurrent, Internal12vCurrent, PowerFactor, SetPhases, MaxPhases, ChargerOfflinePhase, ChargerOfflineCurrent, RcdCalibration, RcdCalibrationNoise, TotalChargePowerSession, SignedMeterValue, SignedMeterValueInterval, SessionEnergyCountExportActive, SessionEnergyCountExportReactive, SessionEnergyCountImportActive, SessionEnergyCountImportReactive, SoftStartTime, ChargeDuration, ChargeMode, ChargePilotLevelInstant, ChargePilotLevelAverage, PilotVsProximityTime, ChargeCurrentInstallationMaxLimit, ChargeCurrentSet, ChargerOperationMode, IsEnabled, IsStandAlone, ChargerCurrentUserUuidDeprecated, CableType, NetworkType, DetectedCar, GridTestResult, FinalStopActive, SessionIdentifier, ChargerCurrentUserUuid, CompletedSession, NewChargeCard, AuthenticationListVersion, EnabledNfcTechnologies, LteRoamingDisabled, InstallationId, RoutingId, Notifications, Warnings, DiagnosticsMode, InternalDiagnosticsLog, DiagnosticsString, CommunicationSignalStrength, CloudConnectionStatus, McuResetSource, McuRxErrors, McuToVariscitePacketErrors, VarisciteToMcuPacketErrors, UptimeVariscite, UptimeMCU, CarSessionLog, CommunicationModeConfigurationInconsistency, RawPilotMonitor, IT3PhaseDiagnosticsLog, PilotTestResults, UnconditionalNfcDetectionIndication, EmcTestCounter, ProductionTestResults, PostProductionTestResults, SmartMainboardSoftwareApplicationVersion, SmartMainboardSoftwareBootloaderVersion, SmartComputerSoftwareApplicationVersion, SmartComputerSoftwareBootloaderVersion, SmartComputerHardwareVersion, MacMain, MacPlcModuleGrid, MacWiFi, MacPlcModuleEv, LteImsi, LteMsisdn, LteIccid, LteImei, MIDCalibration}

var _ObservationIDNameToValueMap = map[string]ObservationID{
	_ObservationIDName[0:15]:           IsOcppConnected,
	_ObservationIDLowerName[0:15]:      IsOcppConnected,
	_ObservationIDName[15:23]:          IsOnline,
	_ObservationIDLowerName[15:23]:     IsOnline,
	_ObservationIDName[23:28]:          Pulse,
	_ObservationIDLowerName[23:28]:     Pulse,
	_ObservationIDName[28:35]:          Unknown,
	_ObservationIDLowerName[28:35]:     Unknown,
	_ObservationIDName[35:46]:          OfflineMode,
	_ObservationIDLowerName[35:46]:     OfflineMode,
	_ObservationIDName[46:68]:          AuthenticationRequired,
	_ObservationIDLowerName[46:68]:     AuthenticationRequired,
	_ObservationIDName[68:81]:          PaymentActive,
	_ObservationIDLowerName[68:81]:     PaymentActive,
	_ObservationIDName[81:96]:          PaymentCurrency,
	_ObservationIDLowerName[81:96]:     PaymentCurrency,
	_ObservationIDName[96:119]:         PaymentSessionUnitPrice,
	_ObservationIDLowerName[96:119]:    PaymentSessionUnitPrice,
	_ObservationIDName[119:141]:        PaymentEnergyUnitPrice,
	_ObservationIDLowerName[119:141]:   PaymentEnergyUnitPrice,
	_ObservationIDName[141:161]:        PaymentTimeUnitPrice,
	_ObservationIDLowerName[141:161]:   PaymentTimeUnitPrice,
	_ObservationIDName[161:178]:        CommunicationMode,
	_ObservationIDLowerName[161:178]:   CommunicationMode,
	_ObservationIDName[178:196]:        PermanentCableLock,
	_ObservationIDLowerName[178:196]:   PermanentCableLock,
	_ObservationIDName[196:207]:        ProductCode,
	_ObservationIDLowerName[196:207]:   ProductCode,
	_ObservationIDName[207:220]:        HmiBrightness,
	_ObservationIDLowerName[207:220]:   HmiBrightness,
	_ObservationIDName[220:242]:        LockCableWhenConnected,
	_ObservationIDLowerName[220:242]:   LockCableWhenConnected,
	_ObservationIDName[242:259]:        SoftStartDisabled,
	_ObservationIDLowerName[242:259]:   SoftStartDisabled,
	_ObservationIDName[259:274]:        FirmwareApiHost,
	_ObservationIDLowerName[259:274]:   FirmwareApiHost,
	_ObservationIDName[274:289]:        MIDBlinkEnabled,
	_ObservationIDLowerName[274:289]:   MIDBlinkEnabled,
	_ObservationIDName[289:309]:        TemperatureInternal5,
	_ObservationIDLowerName[289:309]:   TemperatureInternal5,
	_ObservationIDName[309:329]:        TemperatureInternal6,
	_ObservationIDLowerName[309:329]:   TemperatureInternal6,
	_ObservationIDName[329:353]:        TemperatureInternalLimit,
	_ObservationIDLowerName[329:353]:   TemperatureInternalLimit,
	_ObservationIDName[353:380]:        TemperatureInternalMaxLimit,
	_ObservationIDLowerName[353:380]:   TemperatureInternalMaxLimit,
	_ObservationIDName[380:388]:        Humidity,
	_ObservationIDLowerName[380:388]:   Humidity,
	_ObservationIDName[388:401]:        VoltagePhase1,
	_ObservationIDLowerName[388:401]:   VoltagePhase1,
	_ObservationIDName[401:414]:        VoltagePhase2,
	_ObservationIDLowerName[401:414]:   VoltagePhase2,
	_ObservationIDName[414:427]:        VoltagePhase3,
	_ObservationIDLowerName[414:427]:   VoltagePhase3,
	_ObservationIDName[427:440]:        CurrentPhase1,
	_ObservationIDLowerName[427:440]:   CurrentPhase1,
	_ObservationIDName[440:453]:        CurrentPhase2,
	_ObservationIDLowerName[440:453]:   CurrentPhase2,
	_ObservationIDName[453:466]:        CurrentPhase3,
	_ObservationIDLowerName[453:466]:   CurrentPhase3,
	_ObservationIDName[466:483]:        ChargerMaxCurrent,
	_ObservationIDLowerName[466:483]:   ChargerMaxCurrent,
	_ObservationIDName[483:500]:        ChargerMinCurrent,
	_ObservationIDLowerName[483:500]:   ChargerMinCurrent,
	_ObservationIDName[500:512]:        ActivePhases,
	_ObservationIDLowerName[500:512]:   ActivePhases,
	_ObservationIDName[512:528]:        TotalChargePower,
	_ObservationIDLowerName[512:528]:   TotalChargePower,
	_ObservationIDName[528:538]:        RcdCurrent,
	_ObservationIDLowerName[528:538]:   RcdCurrent,
	_ObservationIDName[538:556]:        Internal12vCurrent,
	_ObservationIDLowerName[538:556]:   Internal12vCurrent,
	_ObservationIDName[556:567]:        PowerFactor,
	_ObservationIDLowerName[556:567]:   PowerFactor,
	_ObservationIDName[567:576]:        SetPhases,
	_ObservationIDLowerName[567:576]:   SetPhases,
	_ObservationIDName[576:585]:        MaxPhases,
	_ObservationIDLowerName[576:585]:   MaxPhases,
	_ObservationIDName[585:604]:        ChargerOfflinePhase,
	_ObservationIDLowerName[585:604]:   ChargerOfflinePhase,
	_ObservationIDName[604:625]:        ChargerOfflineCurrent,
	_ObservationIDLowerName[604:625]:   ChargerOfflineCurrent,
	_ObservationIDName[625:639]:        RcdCalibration,
	_ObservationIDLowerName[625:639]:   RcdCalibration,
	_ObservationIDName[639:658]:        RcdCalibrationNoise,
	_ObservationIDLowerName[639:658]:   RcdCalibrationNoise,
	_ObservationIDName[658:681]:        TotalChargePowerSession,
	_ObservationIDLowerName[658:681]:   TotalChargePowerSession,
	_ObservationIDName[681:697]:        SignedMeterValue,
	_ObservationIDLowerName[681:697]:   SignedMeterValue,
	_ObservationIDName[697:721]:        SignedMeterValueInterval,
	_ObservationIDLowerName[697:721]:   SignedMeterValueInterval,
	_ObservationIDName[721:751]:        SessionEnergyCountExportActive,
	_ObservationIDLowerName[721:751]:   SessionEnergyCountExportActive,
	_ObservationIDName[751:783]:        SessionEnergyCountExportReactive,
	_ObservationIDLowerName[751:783]:   SessionEnergyCountExportReactive,
	_ObservationIDName[783:813]:        SessionEnergyCountImportActive,
	_ObservationIDLowerName[783:813]:   SessionEnergyCountImportActive,
	_ObservationIDName[813:845]:        SessionEnergyCountImportReactive,
	_ObservationIDLowerName[813:845]:   SessionEnergyCountImportReactive,
	_ObservationIDName[845:858]:        SoftStartTime,
	_ObservationIDLowerName[845:858]:   SoftStartTime,
	_ObservationIDName[858:872]:        ChargeDuration,
	_ObservationIDLowerName[858:872]:   ChargeDuration,
	_ObservationIDName[872:882]:        ChargeMode,
	_ObservationIDLowerName[872:882]:   ChargeMode,
	_ObservationIDName[882:905]:        ChargePilotLevelInstant,
	_ObservationIDLowerName[882:905]:   ChargePilotLevelInstant,
	_ObservationIDName[905:928]:        ChargePilotLevelAverage,
	_ObservationIDLowerName[905:928]:   ChargePilotLevelAverage,
	_ObservationIDName[928:948]:        PilotVsProximityTime,
	_ObservationIDLowerName[928:948]:   PilotVsProximityTime,
	_ObservationIDName[948:981]:        ChargeCurrentInstallationMaxLimit,
	_ObservationIDLowerName[948:981]:   ChargeCurrentInstallationMaxLimit,
	_ObservationIDName[981:997]:        ChargeCurrentSet,
	_ObservationIDLowerName[981:997]:   ChargeCurrentSet,
	_ObservationIDName[997:1017]:       ChargerOperationMode,
	_ObservationIDLowerName[997:1017]:  ChargerOperationMode,
	_ObservationIDName[1017:1026]:      IsEnabled,
	_ObservationIDLowerName[1017:1026]: IsEnabled,
	_ObservationIDName[1026:1038]:      IsStandAlone,
	_ObservationIDLowerName[1026:1038]: IsStandAlone,
	_ObservationIDName[1038:1070]:      ChargerCurrentUserUuidDeprecated,
	_ObservationIDLowerName[1038:1070]: ChargerCurrentUserUuidDeprecated,
	_ObservationIDName[1070:1079]:      CableType,
	_ObservationIDLowerName[1070:1079]: CableType,
	_ObservationIDName[1079:1090]:      NetworkType,
	_ObservationIDLowerName[1079:1090]: NetworkType,
	_ObservationIDName[1090:1101]:      DetectedCar,
	_ObservationIDLowerName[1090:1101]: DetectedCar,
	_ObservationIDName[1101:1115]:      GridTestResult,
	_ObservationIDLowerName[1101:1115]: GridTestResult,
	_ObservationIDName[1115:1130]:      FinalStopActive,
	_ObservationIDLowerName[1115:1130]: FinalStopActive,
	_ObservationIDName[1130:1147]:      SessionIdentifier,
	_ObservationIDLowerName[1130:1147]: SessionIdentifier,
	_ObservationIDName[1147:1169]:      ChargerCurrentUserUuid,
	_ObservationIDLowerName[1147:1169]: ChargerCurrentUserUuid,
	_ObservationIDName[1169:1185]:      CompletedSession,
	_ObservationIDLowerName[1169:1185]: CompletedSession,
	_ObservationIDName[1185:1198]:      NewChargeCard,
	_ObservationIDLowerName[1185:1198]: NewChargeCard,
	_ObservationIDName[1198:1223]:      AuthenticationListVersion,
	_ObservationIDLowerName[1198:1223]: AuthenticationListVersion,
	_ObservationIDName[1223:1245]:      EnabledNfcTechnologies,
	_ObservationIDLowerName[1223:1245]: EnabledNfcTechnologies,
	_ObservationIDName[1245:1263]:      LteRoamingDisabled,
	_ObservationIDLowerName[1245:1263]: LteRoamingDisabled,
	_ObservationIDName[1263:1277]:      InstallationId,
	_ObservationIDLowerName[1263:1277]: InstallationId,
	_ObservationIDName[1277:1286]:      RoutingId,
	_ObservationIDLowerName[1277:1286]: RoutingId,
	_ObservationIDName[1286:1299]:      Notifications,
	_ObservationIDLowerName[1286:1299]: Notifications,
	_ObservationIDName[1299:1307]:      Warnings,
	_ObservationIDLowerName[1299:1307]: Warnings,
	_ObservationIDName[1307:1322]:      DiagnosticsMode,
	_ObservationIDLowerName[1307:1322]: DiagnosticsMode,
	_ObservationIDName[1322:1344]:      InternalDiagnosticsLog,
	_ObservationIDLowerName[1322:1344]: InternalDiagnosticsLog,
	_ObservationIDName[1344:1361]:      DiagnosticsString,
	_ObservationIDLowerName[1344:1361]: DiagnosticsString,
	_ObservationIDName[1361:1388]:      CommunicationSignalStrength,
	_ObservationIDLowerName[1361:1388]: CommunicationSignalStrength,
	_ObservationIDName[1388:1409]:      CloudConnectionStatus,
	_ObservationIDLowerName[1388:1409]: CloudConnectionStatus,
	_ObservationIDName[1409:1423]:      McuResetSource,
	_ObservationIDLowerName[1409:1423]: McuResetSource,
	_ObservationIDName[1423:1434]:      McuRxErrors,
	_ObservationIDLowerName[1423:1434]: McuRxErrors,
	_ObservationIDName[1434:1460]:      McuToVariscitePacketErrors,
	_ObservationIDLowerName[1434:1460]: McuToVariscitePacketErrors,
	_ObservationIDName[1460:1486]:      VarisciteToMcuPacketErrors,
	_ObservationIDLowerName[1460:1486]: VarisciteToMcuPacketErrors,
	_ObservationIDName[1486:1501]:      UptimeVariscite,
	_ObservationIDLowerName[1486:1501]: UptimeVariscite,
	_ObservationIDName[1501:1510]:      UptimeMCU,
	_ObservationIDLowerName[1501:1510]: UptimeMCU,
	_ObservationIDName[1510:1523]:      CarSessionLog,
	_ObservationIDLowerName[1510:1523]: CarSessionLog,
	_ObservationIDName[1523:1566]:      CommunicationModeConfigurationInconsistency,
	_ObservationIDLowerName[1523:1566]: CommunicationModeConfigurationInconsistency,
	_ObservationIDName[1566:1581]:      RawPilotMonitor,
	_ObservationIDLowerName[1566:1581]: RawPilotMonitor,
	_ObservationIDName[1581:1603]:      IT3PhaseDiagnosticsLog,
	_ObservationIDLowerName[1581:1603]: IT3PhaseDiagnosticsLog,
	_ObservationIDName[1603:1619]:      PilotTestResults,
	_ObservationIDLowerName[1603:1619]: PilotTestResults,
	_ObservationIDName[1619:1654]:      UnconditionalNfcDetectionIndication,
	_ObservationIDLowerName[1619:1654]: UnconditionalNfcDetectionIndication,
	_ObservationIDName[1654:1668]:      EmcTestCounter,
	_ObservationIDLowerName[1654:1668]: EmcTestCounter,
	_ObservationIDName[1668:1689]:      ProductionTestResults,
	_ObservationIDLowerName[1668:1689]: ProductionTestResults,
	_ObservationIDName[1689:1714]:      PostProductionTestResults,
	_ObservationIDLowerName[1689:1714]: PostProductionTestResults,
	_ObservationIDName[1714:1754]:      SmartMainboardSoftwareApplicationVersion,
	_ObservationIDLowerName[1714:1754]: SmartMainboardSoftwareApplicationVersion,
	_ObservationIDName[1754:1793]:      SmartMainboardSoftwareBootloaderVersion,
	_ObservationIDLowerName[1754:1793]: SmartMainboardSoftwareBootloaderVersion,
	_ObservationIDName[1793:1832]:      SmartComputerSoftwareApplicationVersion,
	_ObservationIDLowerName[1793:1832]: SmartComputerSoftwareApplicationVersion,
	_ObservationIDName[1832:1870]:      SmartComputerSoftwareBootloaderVersion,
	_ObservationIDLowerName[1832:1870]: SmartComputerSoftwareBootloaderVersion,
	_ObservationIDName[1870:1898]:      SmartComputerHardwareVersion,
	_ObservationIDLowerName[1870:1898]: SmartComputerHardwareVersion,
	_ObservationIDName[1898:1905]:      MacMain,
	_ObservationIDLowerName[1898:1905]: MacMain,
	_ObservationIDName[1905:1921]:      MacPlcModuleGrid,
	_ObservationIDLowerName[1905:1921]: MacPlcModuleGrid,
	_ObservationIDName[1921:1928]:      MacWiFi,
	_ObservationIDLowerName[1921:1928]: MacWiFi,
	_ObservationIDName[1928:1942]:      MacPlcModuleEv,
	_ObservationIDLowerName[1928:1942]: MacPlcModuleEv,
	_ObservationIDName[1942:1949]:      LteImsi,
	_ObservationIDLowerName[1942:1949]: LteImsi,
	_ObservationIDName[1949:1958]:      LteMsisdn,
	_ObservationIDLowerName[1949:1958]: LteMsisdn,
	_ObservationIDName[1958:1966]:      LteIccid,
	_ObservationIDLowerName[1958:1966]: LteIccid,
	_ObservationIDName[1966:1973]:      LteImei,
	_ObservationIDLowerName[1966:1973]: LteImei,
	_ObservationIDName[1973:1987]:      MIDCalibration,
	_ObservationIDLowerName[1973:1987]: MIDCalibration,
}

var _ObservationIDNames = []string{
	_ObservationIDName[0:15],
	_ObservationIDName[15:23],
	_ObservationIDName[23:28],
	_ObservationIDName[28:35],
	_ObservationIDName[35:46],
	_ObservationIDName[46:68],
	_ObservationIDName[68:81],
	_ObservationIDName[81:96],
	_ObservationIDName[96:119],
	_ObservationIDName[119:141],
	_ObservationIDName[141:161],
	_ObservationIDName[161:178],
	_ObservationIDName[178:196],
	_ObservationIDName[196:207],
	_ObservationIDName[207:220],
	_ObservationIDName[220:242],
	_ObservationIDName[242:259],
	_ObservationIDName[259:274],
	_ObservationIDName[274:289],
	_ObservationIDName[289:309],
	_ObservationIDName[309:329],
	_ObservationIDName[329:353],
	_ObservationIDName[353:380],
	_ObservationIDName[380:388],
	_ObservationIDName[388:401],
	_ObservationIDName[401:414],
	_ObservationIDName[414:427],
	_ObservationIDName[427:440],
	_ObservationIDName[440:453],
	_ObservationIDName[453:466],
	_ObservationIDName[466:483],
	_ObservationIDName[483:500],
	_ObservationIDName[500:512],
	_ObservationIDName[512:528],
	_ObservationIDName[528:538],
	_ObservationIDName[538:556],
	_ObservationIDName[556:567],
	_ObservationIDName[567:576],
	_ObservationIDName[576:585],
	_ObservationIDName[585:604],
	_ObservationIDName[604:625],
	_ObservationIDName[625:639],
	_ObservationIDName[639:658],
	_ObservationIDName[658:681],
	_ObservationIDName[681:697],
	_ObservationIDName[697:721],
	_ObservationIDName[721:751],
	_ObservationIDName[751:783],
	_ObservationIDName[783:813],
	_ObservationIDName[813:845],
	_ObservationIDName[845:858],
	_ObservationIDName[858:872],
	_ObservationIDName[872:882],
	_ObservationIDName[882:905],
	_ObservationIDName[905:928],
	_ObservationIDName[928:948],
	_ObservationIDName[948:981],
	_ObservationIDName[981:997],
	_ObservationIDName[997:1017],
	_ObservationIDName[1017:1026],
	_ObservationIDName[1026:1038],
	_ObservationIDName[1038:1070],
	_ObservationIDName[1070:1079],
	_ObservationIDName[1079:1090],
	_ObservationIDName[1090:1101],
	_ObservationIDName[1101:1115],
	_ObservationIDName[1115:1130],
	_ObservationIDName[1130:1147],
	_ObservationIDName[1147:1169],
	_ObservationIDName[1169:1185],
	_ObservationIDName[1185:1198],
	_ObservationIDName[1198:1223],
	_ObservationIDName[1223:1245],
	_ObservationIDName[1245:1263],
	_ObservationIDName[1263:1277],
	_ObservationIDName[1277:1286],
	_ObservationIDName[1286:1299],
	_ObservationIDName[1299:1307],
	_ObservationIDName[1307:1322],
	_ObservationIDName[1322:1344],
	_ObservationIDName[1344:1361],
	_ObservationIDName[1361:1388],
	_ObservationIDName[1388:1409],
	_ObservationIDName[1409:1423],
	_ObservationIDName[1423:1434],
	_ObservationIDName[1434:1460],
	_ObservationIDName[1460:1486],
	_ObservationIDName[1486:1501],
	_ObservationIDName[1501:1510],
	_ObservationIDName[1510:1523],
	_ObservationIDName[1523:1566],
	_ObservationIDName[1566:1581],
	_ObservationIDName[1581:1603],
	_ObservationIDName[1603:1619],
	_ObservationIDName[1619:1654],
	_ObservationIDName[1654:1668],
	_ObservationIDName[1668:1689],
	_ObservationIDName[1689:1714],
	_ObservationIDName[1714:1754],
	_ObservationIDName[1754:1793],
	_ObservationIDName[1793:1832],
	_ObservationIDName[1832:1870],
	_ObservationIDName[1870:1898],
	_ObservationIDName[1898:1905],
	_ObservationIDName[1905:1921],
	_ObservationIDName[1921:1928],
	_ObservationIDName[1928:1942],
	_ObservationIDName[1942:1949],
	_ObservationIDName[1949:1958],
	_ObservationIDName[1958:1966],
	_ObservationIDName[1966:1973],
	_ObservationIDName[1973:1987],
}

// ObservationIDString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ObservationIDString(s string) (ObservationID, error) {
	if val, ok := _ObservationIDNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _ObservationIDNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ObservationID values", s)
}

// ObservationIDValues returns all values of the enum
func ObservationIDValues() []ObservationID {
	return _ObservationIDValues
}

// ObservationIDStrings returns a slice of all String values of the enum
func ObservationIDStrings() []string {
	strs := make([]string, len(_ObservationIDNames))
	copy(strs, _ObservationIDNames)
	return strs
}

// IsAObservationID returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ObservationID) IsAObservationID() bool {
	_, ok := _ObservationIDMap[i]
	return ok
}
