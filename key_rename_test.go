package json_your_way

import (
    "regexp"
    "strings"
    "testing"
)

type TestUser struct {
    Id           string `json:"IIIIIDDDDD"`
    FullName     string
    Salutation   string
    FriendlyName string
    Company      string
    Email        string
    Username     string
    Password     string
}

type TestUserGroup struct {
    Id          string
    GroupType   string
    Description string `json:"SillyNamed__Thing"`
    Members     []*TestUser
    Managers    []*TestUser
}

func TestKeyRename(t *testing.T) {
    CustomJsonKey = &StructKeyRenameInterface{
        StructToJson: func(key string) string {
            re := regexp.MustCompile(`([a-z])([A-Z])`)
            return strings.ToLower(re.ReplaceAllString(key, "${1}_${2}"))
        },
    }

    ug := &TestUserGroup{
        Id: "11",
        Description: "DESC 1",
        Members: []*TestUser{
            &TestUser{FullName: "ABCS", FriendlyName: "Friendly Name 1", Email: "b@c.com" },
            &TestUser{},
        },
    }
    out, err := Marshal(ug)
    if err != nil {
        t.Errorf("Custom key encode failed: %s", err)
    }
    ugEncoded := []byte("{\"id\":\"11\",\"group_type\":\"\",\"SillyNamed__Thing\":\"DESC 1\",\"members\":[{\"IIIIIDDDDD\":\"\",\"full_name\":\"ABCS\",\"salutation\":\"\",\"friendly_name\":\"Friendly Name 1\",\"company\":\"\",\"email\":\"b@c.com\",\"username\":\"\",\"password\":\"\"},{\"IIIIIDDDDD\":\"\",\"full_name\":\"\",\"salutation\":\"\",\"friendly_name\":\"\",\"company\":\"\",\"email\":\"\",\"username\":\"\",\"password\":\"\"}],\"managers\":null}")
    if string(out) != string(ugEncoded) {
        t.Errorf("Key rename encoding does not match expected results.\nSource: %s\n Match: %s", ugEncoded, out)
    }

    ug2 := &TestUserGroup{}
    Unmarshal(ugEncoded, ug2)
    if ug2.Members[0].FriendlyName != "Friendly Name 1" {
        t.Errorf("Key rename decoding does not match expected results.")
    }
    CustomJsonKey = nil
}
