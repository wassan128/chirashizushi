package chirashi

import (
    "testing"
)

func TestOpen(t *testing.T) {
    t.Logf("[*] TestOpen")

    shop := Open("1")

    actual_Name := shop.Name
    expected_Name := "フーディアム 堂島"
    if actual_Name != expected_Name {
        t.Errorf("got: %v\nwant: %v\n", actual_Name, expected_Name)
    }

    actual_Id := shop.Id
    expected_Id := "1"
    if actual_Id != expected_Id {
        t.Errorf("got: %v\nwant: %v\n", actual_Id, expected_Id)
    }
}

