Limit range extension - initial notes
Concept has to be discussed and can/will be changed

#Changes applied to API:

``` golang
+type LimitRangeConstraint string
+
+const (
+       CPURequest LimitRangeConstraint = "requests.cpu"
+       MemoryRequest LimitRangeConstraint = "requests.memory"
+       CPULimit LimitRangeConstraint = "limits.cpu"
+       MemoryLimit LimitRangeConstraint = "limits.memory"
+       CPUBurstable LimitRangeConstraint = "burstable.cpu"
+       MemoryBurstable LimitRangeConstraint = "burstable.memory"
+       StorageRequest LimitRangeConstraint = "requests.storage"
+)
+ 
+ type ResourceConstraintsList map[LimitRangeConstraint]resource.Quantity

type LimitRangeItem struct {
        Type LimitType `json:"type,omitempty"`
        // Max usage constraints on this kind by resource name
        // +optional
-       Max ResourceList `json:"max,omitempty"`
+       Max ResourceConstraintsList `json:"max,omitempty"`
        // Min usage constraints on this kind by resource name
        // +optional
-       Min ResourceList `json:"min,omitempty"`
+       Min LimitRangeConstraint `json:"min,omitempty"`
        // Default resource requirement limit value by resource name.
        // +optional
-        Default ResourceList `json:"default,omitempty"`
+       Default LimitRangeConstraint `json:"min,omitempty"`
-        // DefaultRequest resource requirement request value by resource name.
-        // +optional
-        DefaultRequest ResourceList `json:"defaultRequest,omitempty"`
-        // MaxLimitRequestRatio represents the max burst value for the named resource
-        // +optional
-        MaxLimitRequestRatio ResourceList `json:"maxLimitRequestRatio,omitempty"`
 }
```

#LimitRangeItem validation:
For for all resources' requests and limits it is true that:

```
Min <= default <= max
```

```
max request <= min limit - maybe this is too strong...??
```

For for all resources' burstable fields is true that:

```
min < max - (or min <= max??)
```

Default.BurstableMemory has no purpose and shouldn't be definable

For PVC - limit and burstable fields are not defined.

#Rules for setting default values of limitRangeConstraints

##LimitType container
Default limit:
```
If !default.limit & max.limit => default.limit = max.limit
If !default.limit & !max.limit & min.limit => default.limit = min.limit
```

Default request:
```
if !default.request & min.request =>  default.request = min.request
if !default.request & !min.request & max.request => default.request = max.request
```
```
if !default.request & !min.request & !max.request & default.limit => default.request = default.limit
```

 - Here we preffer requests values over default.limit unlike how it is done in current concept - is this ok??
 
##LimitType pod
 - no defaults
  
##LimitType PVC
 - no defaults

#Limit range admission rules

For containers and pods

```
( limitrangeitem.burstable.max is defined => container.limit is defined & container.request is defined) &
    ( min.burstable(or 0) <= container.limit/container.request <= max.burstable)
```

```
limitrangeitem.min.limit <= container.limit <= limitrangeitem.max.limit
```

For containers, pods and PVC
```
limitrangeitem.min.request <= container.limit <= limitrangeitem.max.request
```
