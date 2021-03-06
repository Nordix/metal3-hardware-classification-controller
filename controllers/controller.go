package controllers

// RBAC Rules for HardwareClassification resources
//
// +kubebuilder:rbac:groups=metal3.io,resources=hardwareclassifications,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=metal3.io,resources=hardwareclassifications/status,verbs=get;update;patch

// RBAC rules for BareMetalHost resources
//
// +kubebuilder:rbac:groups=metal3.io,resources=baremetalhosts,verbs=get;list;watch;update
// +kubebuilder:rbac:groups=metal3.io,resources=baremetalhosts/status,verbs=get
