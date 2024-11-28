// internal/resources/constants.go
package resources

// WEEKDAY represents days of the week
type WEEKDAY string

// WEEKDAY constants
const (
	MONDAY    WEEKDAY = "MONDAY"
	TUESDAY   WEEKDAY = "TUESDAY"
	WEDNESDAY WEEKDAY = "WEDNESDAY"
	THURSDAY  WEEKDAY = "THURSDAY"
	FRIDAY    WEEKDAY = "FRIDAY"
	SATURDAY  WEEKDAY = "SATURDAY"
	SUNDAY    WEEKDAY = "SUNDAY"
)

// GENDER represents user gender
type GENDER string

// GENDER constants
const (
	MALE   GENDER = "MALE"
	FEMALE GENDER = "FEMALE"
	OTHER  GENDER = "OTHER"
)

// USERROLE represents different user roles
type USERROLE int32

// USERROLE constants
const (
	SUPERADMIN USERROLE = 1
	ADMIN      USERROLE = 2
	GYM        USERROLE = 3
	GYMSTAFF   USERROLE = 4
	CUSTOMER   USERROLE = 5
	INVALID    USERROLE = 0
)

