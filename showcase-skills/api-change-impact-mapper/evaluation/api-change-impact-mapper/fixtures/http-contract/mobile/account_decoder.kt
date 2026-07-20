data class Account(val id: String, val displayName: String, val status: String)
fun decodeAccount(json: JsonObject) = Account(json["id"]!!.jsonPrimitive.content, json["display_name"]!!.jsonPrimitive.content, json["status"]!!.jsonPrimitive.content)
