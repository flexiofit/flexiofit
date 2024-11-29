// internal/resources/extractor.go
package resources

import (
	"strings"
)

func GetUserType(userType string) USERROLE {
    switch strings.ToUpper(userType) {
    case "SUPERADMIN":
        return SUPERADMIN
    case "ADMIN":
        return ADMIN
    case "GYM":
        return GYM
    case "GYMSTAFF":
        return GYMSTAFF
    case "CUSTOMER":
        return CUSTOMER
    default:
        return INVALID
    }
}

var userRoleNames = map[USERROLE]string{
    INVALID:    "INVALID",
    SUPERADMIN: "SUPERADMIN",
    ADMIN:      "ADMIN",
    GYM:        "GYM",
    GYMSTAFF:   "GYMSTAFF",
    CUSTOMER:   "CUSTOMER",
}

func (r USERROLE) String() string {
    if name, exists := userRoleNames[r]; exists {
        return name
    }
    return "UNKNOWN"
}
