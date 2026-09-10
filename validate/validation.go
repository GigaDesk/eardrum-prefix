package validate

import (
	"encoding/base64"
	pgerror "errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/GigaDesk/eardrum-interfaces/errors"
)

const (
	// 1. Updated dimensions to match your MobileFaceNet model
	MobileFaceNetDim = 128
	BytesPerFloat32  = 4
	ExpectedByteLen  = MobileFaceNetDim * BytesPerFloat32 // 512 bytes
)

// ValidateMobileFaceNetEmbeddingB64 strictly validates a MobileFaceNet base64 string.
// It checks base64 compliance, length, and optional L2 normalization.
func ValidateMobileFaceNetEmbeddingB64(b64Str string, checkNormalization bool) (bool, error) {
	// 1. Strict base64 decoding
	decodedBytes, err := base64.StdEncoding.Strict().DecodeString(b64Str)
	if err != nil {
		err1 := errors.New(errors.EAREmbeddingInvalidBase64, err)
		err1.Log()
		return false, err1
	}

	// 2. Exact size check for MobileFaceNet (128 float32s = 512 bytes)
	if len(decodedBytes) != ExpectedByteLen {
		err1 := errors.New(errors.EAREmbeddingInvalidLength, pgerror.New("invalid byte length; MobileFaceNet embeddings must be exactly 512 bytes (128 float32s)"))
		err1.Log()
		return false, err1
	}

	// 3. Optional: Mathematical L2 Normalization Check
	if checkNormalization {
		var sumOfSquares float32 = 0.0

		// Read bytes directly as float32 to calculate magnitude
		for i := 0; i < ExpectedByteLen; i += BytesPerFloat32 {
			// Fast bit-reconstruction for float32 (assuming standard Little Endian data)
			bits := uint32(decodedBytes[i]) |
				uint32(decodedBytes[i+1])<<8 |
				uint32(decodedBytes[i+2])<<16 |
				uint32(decodedBytes[i+3])<<24

			val := math.Float32frombits(bits)

			// Guard against NaN or Infinity values leaking through
			if math.IsNaN(float64(val)) || math.IsInf(float64(val), 0) {
				err1 := errors.New(errors.EAREmbeddingInvalidFloatData, pgerror.New("embedding contains invalid float data (NaN or Inf)"))
				err1.Log()
				return false, err1
			}

			sumOfSquares += val * val
		}

		// The magnitude should equal 1.0. We allow a tiny epsilon tolerance for float rounding.
		magnitude := math.Sqrt(float64(sumOfSquares))
		const epsilon = 1e-3
		if math.Abs(magnitude-1.0) > epsilon {
			err1 := errors.New(errors.EAREmbeddingNotNormalized, pgerror.New("embedding is not L2 normalized"))
			err1.Log()
			return false, err1
		}
	}

	return true, nil
}

/*
ValidatePhoneNumber partially checks the validity of a phone number

Warning: While ValidatePhoneNumber instills some guard rails it does not guarantee validity of a phone number but only checks for basic stuff like country code format
*/
func ValidatePhoneNumber(phone_number string) error {
	if phone_number == "" {
		err1 := errors.New(errors.EARValidationInvalidPhone, pgerror.New("phone number is empty"))
		err1.Log()
		return err1
	}
	regex := `^[+]{1}(?:[0-9\-\(\)\/\.\s?]){6,15}[0-9]{1}$`

	matched, err := regexp.MatchString(regex, phone_number)
	if err != nil {
		err1 := errors.New(errors.EARValidationInvalidPhone, err)
		err1.Log()
		return err1
	}
	if !matched {
		err1 := errors.New(errors.EARValidationInvalidPhone, pgerror.New("Invalid phone number"))
		err1.Log()
		return err1
	}
	return nil
}

/*
ValidateKenyanPhoneNumber partially checks the validity of a Kenyan phone number. It is more accurate than ValidatePhoneNumber on Kenyan phone numbers

Warning: While ValidateKenyanPhoneNumber instills some guard rails, and is more accurate than ValidatePhoneNumber, it still does not guarantee the validity of a phone number.
*/
func ValidateKenyanPhoneNumber(phone_number string) error {
	if len(phone_number) != 13 {
		err1 := errors.New(errors.EARValidationInvalidPhone, pgerror.New("Invalid phone number"))
		err1.Log()
		return err1
	}
	err := ValidatePhoneNumber(phone_number)
	if err != nil {
		return err
	}
	return nil
}

// Checks the validity of an OTP code entry
func ValidateOtp(otp string) error {
	if len(otp) < 6 {
		err1 := errors.New(errors.EARValidationInvalidPhone, pgerror.New("Otp is less than six digits"))
		err1.Log()
		return err1
	}
	if len(otp) > 6 {
		err1 := errors.New(errors.EARValidationInvalidPhone, pgerror.New("Otp is more than six digits"))
		err1.Log()
		return err1
	}
	return nil
}

// Checks the validity of a name entry
func ValidateName(name string) error {
	if name == "" {
		err1 := errors.New(errors.EARValidationInvalidName, pgerror.New("name is empty"))
		err1.Log()
		return err1
	}
	return nil
}

// Checks the validity of a password entry
func ValidatePassword(password string) error {
	if len(password) < 8 {
		err1 := errors.New(errors.EARValidationInvalidPassword, pgerror.New("password must be atleast 8 characters"))
		err1.Log()
		return err1
	}
	return nil
}

/*
Checks the validity of a date entry in the format:

	yyyy-mm-dd
*/
func ValidateDate(date string) error {
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		err1 := errors.New(errors.EARValidationInvalidDateFormat, err)
		err1.Log()
		return err1
	}
	return nil
}

/*
Checks the validity of a date that is in the past or current. Throws an error if the date provided is in the future
*/
func ValidatePastDate(date time.Time) error {
	if date.After(time.Now()) {
		err1 := errors.New(errors.EARValidationFutureDateNotAllowed, pgerror.New("cannot accept future date"))
		err1.Log()
		return err1
	}
	return nil
}

// ValidateOfflineTransactionID checks if the provided offline transaction ID
// adheres to the `IMEI_MillisecondTimestamp` format, verifying the Luhn checksum
// of the IMEI and the plausibility of the millisecond timestamp.
func ValidateOfflineTransactionID(id string) error {
	if strings.TrimSpace(id) == "" {
		err1 := errors.New(errors.EARValidationInvalidOfflineTxId, pgerror.New("offline transaction ID cannot be empty"))
		err1.Log()
		return err1
	}

	parts := strings.Split(id, "_")
	if len(parts) != 2 {
		err1 := errors.New(errors.EARValidationInvalidOfflineTxId, pgerror.New("invalid offline transaction ID format: expected 'IMEI_Timestamp'"))
		err1.Log()
		return err1
	}

	imei := parts[0]
	timestampStr := parts[1]

	// 1. Validate IMEI format & Luhn checksum
	if err := validateIMEI(imei); err != nil {
		err1 := errors.New(errors.EARValidationInvalidIMEI, err)
		err1.Log()
		return err1
	}

	// 2. Validate Timestamp Structure (Must be 13 digits for millisecond Epoch)
	if len(timestampStr) != 13 {
		err1 := errors.New(errors.EARValidationInvalidUnixTimestamp, pgerror.New("invalid timestamp: must be a 13-digit millisecond Unix timestamp"))
		err1.Log()
		return err1
	}

	millis, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		err1 := errors.New(errors.EARValidationInvalidUnixTimestamp, fmt.Errorf("failed to parse numeric value: %w", err))
		err1.Log()
		return err1
	}

	// 3. Plausibility Check (Timestamp must not be in the future or prior to 2020)
	const minTimestamp = 1577836800000 // 2020-01-01T00:00:00Z in milliseconds
	nowMillis := time.Now().UnixMilli()

	if millis < minTimestamp {
		err1 := errors.New(errors.EARValidationInvalidUnixTimestamp, pgerror.New("invalid timestamp: date is too far in the past"))
		err1.Log()
		return err1
	}
	if millis > nowMillis+60000 { // Allow 1-minute buffer for minor client clock drift
		err1 := errors.New(errors.EARValidationInvalidUnixTimestamp, pgerror.New("invalid timestamp: date cannot be in the future"))
		err1.Log()
		return err1
	}

	return nil
}

// validateIMEI checks if the IMEI is exactly 15 numeric digits and passes the Luhn checksum.
func validateIMEI(imei string) error {
	if len(imei) != 15 {
		return pgerror.New("must be exactly 15 digits long")
	}

	var sum int
	for i := 0; i < 15; i++ {
		digit, err := strconv.Atoi(string(imei[i]))
		if err != nil {
			return pgerror.New("must contain only numeric digits")
		}

		// Double every second digit (at odd indices 1, 3, 5...)
		if i%2 == 1 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}

	if sum%10 != 0 {
		return pgerror.New("failed Luhn checksum validation")
	}

	return nil
}
