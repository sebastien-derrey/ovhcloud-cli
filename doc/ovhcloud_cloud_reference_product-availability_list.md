## ovhcloud cloud reference product-availability list

List the available products, one row per product type, model and location

```
ovhcloud cloud reference product-availability list [flags]
```

### Options

```
      --filter stringArray     Filter results by any property using https://github.com/PaesslerAG/gval syntax
                               Examples:
                                 --filter 'state=="running"'
                                 --filter 'name=~"^my.*"'
                                 --filter 'nested.property.subproperty>10'
                                 --filter 'startDate>="2023-12-01"'
                                 --filter 'name=~"something" && nbField>10'
  -h, --help                   help for list
      --name string            Exact model name (e.g., b3-8)
      --name-contains string   Part of the model name, letter case ignored
  -r, --region string          Region, letter case ignored (e.g., GRA11, EU-WEST-PAR)
      --type string            Product type (FLOATING_IP, GATEWAY, INSTANCE, INSTANCE_SNAPSHOT, LOADBALANCER, OBJECT_STORAGE, PUBLIC_IP, SHARE, SHARE_SNAPSHOT, SWIFT, VOLUME, VOLUME_BACKUP, VOLUME_SNAPSHOT)
```

### Options inherited from parent commands

```
      --cloud-project string   Cloud project ID
  -d, --debug                  Activate debug mode (will log all HTTP requests details)
  -e, --ignore-errors          Ignore errors in API calls when it is not fatal to the execution
  -o, --output string          Output format: json, yaml, interactive, or a custom format expression (using https://github.com/PaesslerAG/gval syntax)
                               Examples:
                                 --output json
                                 --output yaml
                                 --output interactive
                                 --output 'id' (to extract a single field)
                                 --output 'nested.field.subfield' (to extract a nested field)
                                 --output '[id, "name"]' (to extract multiple fields as an array)
                                 --output '{"newKey": oldKey, "otherKey": nested.field}' (to extract and rename fields in an object)
                                 --output 'name+","+type' (to extract and concatenate fields in a string)
                                 --output '(nbFieldA + nbFieldB) * 10' (to compute values from numeric fields)
      --profile string         Use a specific profile from the configuration file
```

### SEE ALSO

* [ovhcloud cloud reference product-availability](ovhcloud_cloud_reference_product-availability.md)	 - Check the products available in the given cloud project

