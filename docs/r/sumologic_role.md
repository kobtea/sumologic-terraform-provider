# sumologic_role
Provides a [Sumologic Role][1].

## Example Usage
```hcl
resource "sumologic_role" "example_role" {
  name        = "TestRole123"
  description = "Testing resource sumologic_role"

  filter_predicate = "_sourceCategory=Test"
  capabilities = [
    "manageCollectors"
  ]
}
```

## Argument reference
The following arguments are supported:
- `name` - (Required) The name of the role.
- `description` - (Optional) The description of the role.
- `filter_predicate` - (Optional) A search filter to restrict access to specific logs.
- `capabilities` - (Optional) List of capabilities associated with this role.
- `lookup_by_name` - (Optional) Configures an existent role using the same 'name' or creates a new one if non existent. Defaults to false.
- `destroy` - (Optional) Whether or not to delete the role in Sumo when it is removed from Terraform.  Defaults to true.

## Attributes reference
The following attributes are exported:
- `id` - The internal ID of the role.


[Back to Index][0]

[0]: ../README.md
[1]: https://help.sumologic.com/Manage/Users-and-Roles/Manage-Roles
