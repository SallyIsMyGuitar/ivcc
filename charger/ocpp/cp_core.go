package ocpp

import (
	"time"

	"github.com/lorenzodonini/ocpp-go/ocpp1.6/core"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/firmware"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/types"
)

const (
	messageExpiry     = 30 * time.Second
	transactionExpiry = time.Hour
)

func (cp *CP) Authorize(request *core.AuthorizeRequest) (*core.AuthorizeConfirmation, error) {
	cp.log.TRACE.Printf("%T: %+v", request, request)

	// TODO check if this authorizes foreign RFID tags
	res := &core.AuthorizeConfirmation{
		IdTagInfo: &types.IdTagInfo{
			Status: types.AuthorizationStatusAccepted,
		},
	}

	return res, nil
}

func (cp *CP) BootNotification(request *core.BootNotificationRequest) (*core.BootNotificationConfirmation, error) {
	cp.log.TRACE.Printf("%T: %+v", request, request)

	cp.mu.Lock()
	defer cp.mu.Unlock()

	// // required fields by spec
	// cp.Details.Manufacturer = request.ChargePointVendor
	// cp.Details.Model = request.ChargePointModel
	//
	// // optional fields by spec
	// if request.ChargePointSerialNumber != "" {
	// 	cp.Details.SerialNumber = request.ChargePointSerialNumber
	// }
	//
	// if request.FirmwareVersion != "" {
	// 	cp.Details.FirmwareVersion = request.FirmwareVersion
	// }
	// if request.Imsi != "" {
	// 	cp.Details.Imsi = request.Imsi
	// }
	// if request.Iccid != "" {
	// 	cp.Details.Iccid = request.Iccid
	// }
	//
	// if request.MeterType != "" {
	// 	cp.Details.MeterModel = request.MeterType
	// }
	// if request.MeterSerialNumber != "" {
	// 	cp.Details.MeterSerialNumber = request.MeterSerialNumber
	// }

	res := &core.BootNotificationConfirmation{
		CurrentTime: types.NewDateTime(time.Now()),
		Interval:    60, // TODO
		Status:      core.RegistrationStatusAccepted,
	}

	close(cp.bootC)

	return res, nil
}

// timestampValid returns false if status timestamps are outdated
func (cp *CP) timestampValid(t time.Time) bool {
	// reject if expired
	if time.Since(t) > messageExpiry {
		return false
	}

	// assume having a timestamp is better than not
	if cp.status.Timestamp == nil {
		return true
	}

	// reject older values than we already have
	return !t.Before(cp.status.Timestamp.Time)
}

func (cp *CP) StatusNotification(request *core.StatusNotificationRequest) (*core.StatusNotificationConfirmation, error) {
	cp.log.TRACE.Printf("%T: %+v", request, request)

	if request != nil {
		cp.mu.Lock()
		defer cp.mu.Unlock()

		if cp.status == nil {
			cp.status = request
			close(cp.statusC) // signal initial status received
		} else if request.Timestamp == nil || cp.timestampValid(request.Timestamp.Time) {
			cp.status = request
		} else {
			cp.log.TRACE.Printf("ignoring status: %s < %s", request.Timestamp.Time, cp.status.Timestamp)
		}
	}

	return new(core.StatusNotificationConfirmation), nil
}

func (cp *CP) DataTransfer(request *core.DataTransferRequest) (*core.DataTransferConfirmation, error) {
	cp.log.TRACE.Printf("%T: %+v", request, request)

	res := &core.DataTransferConfirmation{
		Status: core.DataTransferStatusRejected,
	}

	return res, nil
}

func (cp *CP) update() {
	cp.mu.Lock()
	cp.updated = time.Now()
	cp.mu.Unlock()
}

func (cp *CP) Heartbeat(request *core.HeartbeatRequest) (*core.HeartbeatConfirmation, error) {
	cp.log.TRACE.Printf("%T: %+v", request, request)

	cp.update()
	res := &core.HeartbeatConfirmation{
		CurrentTime: types.NewDateTime(time.Now()),
	}

	return res, nil
}

func (cp *CP) MeterValues(request *core.MeterValuesRequest) (*core.MeterValuesConfirmation, error) {
	cp.log.TRACE.Printf("%T: %+v", request, request)

	cp.mu.Lock()
	defer cp.mu.Unlock()

	for _, meterValue := range request.MeterValue {
		// ignore old meter value requests
		if meterValue.Timestamp.Time.After(cp.meterUpdated) {
			for _, sample := range meterValue.SampledValue {
				cp.measurements[getSampleKey(sample)] = sample
				cp.meterUpdated = time.Now()
			}
		}
	}

	return new(core.MeterValuesConfirmation), nil
}

func getSampleKey(s types.SampledValue) string {
	if s.Phase != "" {
		return string(s.Measurand) + "@" + string(s.Phase)
	}

	return string(s.Measurand)
}

func (cp *CP) StartTransaction(request *core.StartTransactionRequest) (*core.StartTransactionConfirmation, error) {
	cp.log.TRACE.Printf("%T: %+v", request, request)

	cp.mu.Lock()
	defer cp.mu.Unlock()

	if request == nil { // what!?
		return nil, nil
	}

	res := &core.StartTransactionConfirmation{
		IdTagInfo: &types.IdTagInfo{
			Status: types.AuthorizationStatusAccepted, // accept
		},
		TransactionId: 1, // default
	}

	if (time.Since(request.Timestamp.Time) < transactionExpiry) || // only respect transactions in the last hour, invalidate it
	(cp.currentTransaction != nil && cp.GetTransactionStatus() != RequestedStart) { // we didnt start this transaction, invalidate it
		cp.log.DEBUG.Printf("responded with authorization status blocked for unknown transaction")
		res.IdTagInfo.Status = types.AuthorizationStatusConcurrentTx
	} else { // create new transaction
		res.TransactionId = cp.TransactionID()
	}

	return res, nil
}

func (cp *CP) StopTransaction(request *core.StopTransactionRequest) (*core.StopTransactionConfirmation, error) {
	cp.log.TRACE.Printf("%T: %+v", request, request)

	cp.mu.Lock()
	defer cp.mu.Unlock()

	// signal that transaction has ended
	if cp.currentTransaction != nil {
		if cp.currentTransaction.Status != Running || cp.currentTransaction.Status != Suspended {
			cp.log.ERROR.Printf("stop transaction: stopping not started transaction")
		}

		// log mismatching id because we close transaction anyway
		if request.TransactionId != cp.currentTransaction.Id {
			cp.log.ERROR.Printf("stop transaction: mismatched id %d expected %d", request.TransactionId, cp.currentTransaction.Id)
		}

		cp.SetTransactionStatus(Finished)
	}

	res := &core.StopTransactionConfirmation{
		IdTagInfo: &types.IdTagInfo{
			Status: types.AuthorizationStatusAccepted, // accept
		},
	}

	return res, nil
}

func (cp *CP) DiagnosticStatusNotification(request *firmware.DiagnosticsStatusNotificationRequest) (*firmware.DiagnosticsStatusNotificationConfirmation, error) {
	cp.log.TRACE.Printf("%T: %+v", request, request)

	return &firmware.DiagnosticsStatusNotificationConfirmation{}, nil
}

func (cp *CP) FirmwareStatusNotification(request *firmware.FirmwareStatusNotificationRequest) (*firmware.FirmwareStatusNotificationConfirmation, error) {
	cp.log.TRACE.Printf("%T: %+v", request, request)

	return &firmware.FirmwareStatusNotificationConfirmation{}, nil
}
