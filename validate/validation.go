package validate

import (
	"encoding/base64"
	pgerror "errors"
	"math"
	"regexp"
	"time"

	"github.com/GigaDesk/eardrum-interfaces/errors"
)

const (
	ArcFaceDim      = 512
	BytesPerFloat32 = 4
	ExpectedByteLen = ArcFaceDim * BytesPerFloat32 // 2048 bytes
)

// ValidateArcFaceEmbeddingB64 strictly validates an ArcFace base64 string.
// It checks base64 compliance, length, and optional L2 normalization.
func ValidateArcFaceEmbeddingB64(b64Str string, checkNormalization bool) (bool, error) {
	// 1. Strict base64 decoding
	decodedBytes, err := base64.StdEncoding.Strict().DecodeString(b64Str)
	if err != nil {
		err1 := errors.New(errors.EAREmbeddingInvalidBase64, err)
		err1.Log()
		return false, err1
	}

	// 2. Exact size check for ArcFace (512 float32s = 2048 bytes)
	if len(decodedBytes) != ExpectedByteLen {
		err1 := errors.New(errors.EAREmbeddingInvalidLength, pgerror.New("invalid byte length; ArcFace embeddings must be exactly 2048 bytes (512 float32s)"))
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
			err1 := errors.New(errors.EAREmbeddingNotNormalized, pgerror.New("embedding is not L2 normalized (violates ArcFace structure)"))
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
