package models 

import (
    // "github.com/go-gorp/gorp"
    //"regexp"
    //"time"
)

type Bundles struct {
	BundleId 			    int64 	`db:"bundleid"`
	BundleName		        string 	`db:"bundlename"`	// unique
}

func CreateDefaultBundle(_BundleName string) Bundles {
    bundle_dummy := Bundles {
        BundleId : 0,
        BundleName: _BundleName,
    }
    return bundle_dummy
}