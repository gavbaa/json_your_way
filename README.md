# json_your_way
Exporting Go structs with JSON fields the way you want them named has never been easier!

Are you tired of adding `json:field_name_here` to your 4000 line long Go structs?

Do you detest writing custom MarshalJSON/UnmarshalJSON methods for every struct just to get a different layout?

Look no further!  You can take an entire project and serialize it the way you want to!  Let the code speak for itself:

{{{
type TestUser struct {
    Id           string `json:"CUSTOM_ID_NAME_PRESERVED"`
    FullName     string
    FriendlyName string
    Email        string
}

json_your_way.CustomJsonKey = &StructKeyRenameInterface{
    StructToJson: func(key string) string {
        re := regexp.MustCompile(`([a-z])([A-Z])`)
        return strings.ToLower(re.ReplaceAllString(key, "${1}_${2}"))
    },
}

u := &User{}
json_your_way.MarshalIndent(u, "    ", "")
/* Outputs:
{
    "CUSTOM_ID_NAME_PRESERVED": "",
    "full_name": "",
    "friendly_name": "",
    "email": ""
}
}}}

Unmarshaling works without any further additions, the method that defines how structs field names are converted to JSON keys is used to handle the mapping in both directions.

Methods are compatible with Go's standard JSON methods.  The code is a modification to Go's encoding/json library.
