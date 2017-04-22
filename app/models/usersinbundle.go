package models 

import (
    // "github.com/go-gorp/gorp"
    //"regexp"
    //"time"
)

type UsersInBundle struct {
	BundleId 			    int64 	`db:"bundleid"`
	UserId	    	        int64 	`db:"userid"`	// unique
}

func CreateDefaultUserInBundle(_BundleId int64, _UserId int64) UsersInBundle {
    UserInBundle_dummy := UsersInBundle {
        BundleId : _BundleId,
        UserId : _UserId,
    }
    return UserInBundle_dummy
}