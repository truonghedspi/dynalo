package authz

import data.allow_ids

allow {
    input.path == ["users"]
    input.method == "POST"
}

allow {
    input.identity = "admin"
}

allow {
    input.method = "GET"
}

allow {
    contains(allow_ids, input.user_id) 
}

# Checks if `list` includes an element matching `item`.
contains(list, item) {
    list[_] = item
}