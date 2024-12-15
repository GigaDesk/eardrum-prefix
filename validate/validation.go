package validate

import (
	"errors"
	"net/url"
	"regexp"
	"time"
)

/* 
ValidatePhoneNumber partially checks the validity of a phone number

Warning: While ValidatePhoneNumber instills some guard rails it does not guarantee validity of a phone number but only checks for basic stuff like country code format
*/
func ValidatePhoneNumber(phone_number string) error {
	if phone_number == "" {
		return errors.New("phone number is empty")
	}
	regex := `^[+]{1}(?:[0-9\-\(\)\/\.\s?]){6,15}[0-9]{1}$`

	matched, err := regexp.MatchString(regex, phone_number)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("Invalid phone number")
	}
	return nil
}
/* 
ValidateKenyanPhoneNumber partially checks the validity of a Kenyan phone number. It is more accurate than ValidatePhoneNumber on Kenyan phone numbers

Warning: While ValidateKenyanPhoneNumber instills some guard rails, and is more accurate than ValidatePhoneNumber, it still does not guarantee the validity of a phone number.
*/
func ValidateKenyanPhoneNumber(phone_number string) error {
	if len(phone_number) != 13 {
		return errors.New("Invalid phone number")
	}
	err := ValidatePhoneNumber(phone_number)
	if err != nil {
		return err
	}
	return nil
}

// Checks the validity of an id entry
func ValidateId(id int) error {
	if id < 1 {
		return errors.New("id cannot be a zero value or less")
	}
	return nil
}

//Checks the validity of an OTP code entry
func ValidateOtp(otp string) error {
	if len(otp) < 6 {
		return errors.New("Otp is less than six digits")
	}
	if len(otp) > 6 {
		return errors.New("Otp is more than six digits")
	}
	return nil
}

//Checks the validity of a registration number entry
func ValidateRegistrationNumber(registration_number string) error {
	if registration_number == "" {
		return errors.New("registration number is empty")
	}
	return nil
}

//Checks the validity of a name entry
func ValidateName(name string) error {
	if name == "" {
		return errors.New("name is empty")
	}
	return nil
}

//Checks the validity of a password entry
func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be atleast 8 characters")
	}
	return nil
}

/*Checks the validity of a date entry in the format:  

   yyyy-mm-dd
*/
func ValidateDate(date string) error {
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return err
	}
	return nil
}

/*Checks the validity of a date that is in the past or current. Throws an error if the date provided is in the future

*/
func ValidatePastDate(date time.Time) error {
	if date.After(time.Now()) {
		return errors.New("cannot accept future date")
	}
	return nil
}

// Checks the validity  of a date of admission entry 
func ValidateDateOfAdmission(date_of_admission time.Time) error {
	err := ValidatePastDate(date_of_admission)
	if err != nil {
		return err
	}
	return nil
}

// Checks the validity of a date of birth entry
func ValidateDateOfBirth(date_of_birth time.Time) error {
	err := ValidatePastDate(date_of_birth)
	if err != nil {
		return err
	}
	return nil
}

// Checks the validity of a url entry
func ValidateUrl(s string) error {
	// Parse the URL string
	u, err := url.Parse(s)
	if err != nil {
		return err
	}

	// Check if the scheme is valid (e.g., http, https)
	if u.Scheme != "http" && u.Scheme != "https" {
		return errors.New("url scheme is invalid")
	}

	// Check if the host is valid
	if u.Host == "" {
		return errors.New("url host is invalid")
	}
	return nil
}

//Checks the validity of a profile picture url entry
func ValidateProfilePicture(profile_picture string) error {
	// Parse the URL string
	err := ValidateUrl(profile_picture)
	if err != nil {
		return err
	}
	return nil
}

//Checks the validity of a badge picture url entry
func ValidateBadge(badge string) error {
	// Parse the URL string
	err := ValidateUrl(badge)
	if err != nil {
		return err
	}
	return nil
}

//Checks the validity of a website url entry
func ValidateWebsite(website string) error {
	// Parse the URL string
	err := ValidateUrl(website)
	if err != nil {
		return err
	}
	return nil
}
