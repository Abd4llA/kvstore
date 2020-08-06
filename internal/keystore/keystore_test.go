package keystore

import (
	"testing"
)

const (
	TEST_KEY  = "TEST_KEY_1"
	TEST_VAL  = "TEST_VAL_1"
	SUCCEEDED = "\u2713"
	FAILED    = "\u2717"
)

func TestPutWithStringValue(t *testing.T) {
	t.Log("Given a keystore object")
	{
		k := NewKeyStore()
		t.Logf("\tTest 0:\tWhen inserting %s => %s", TEST_KEY, TEST_VAL)
		{
			if err := k.Put(TEST_KEY, TEST_VAL); err != nil {
				t.Fatalf("\t%s\tShould be put successfully : %v", FAILED, err)
			}
			t.Logf("\t%s\tShould be put successfully.", SUCCEEDED)
		}
		t.Logf("\tTest 1:\tWhen getting value for key %s", TEST_KEY)
		{
			v := k.Get(TEST_KEY)
			if v == "" {
				t.Fatalf("\t%s\tShould return value (%s) succesfully. No value found for key %s", FAILED, TEST_VAL, TEST_KEY)
			}
			if v != TEST_VAL {
				t.Fatalf("\t%s\tShould return value (%s) succesfully. Unexpected value. Received: %s Expected: %s", FAILED, TEST_VAL, v, TEST_VAL)
			}
			t.Logf("\t%s\tShould return value %s succesfully", SUCCEEDED, TEST_VAL)
		}
		t.Logf("\tTest 2:\tWhen deleting %s", TEST_KEY)
		{
			if err := k.Del(TEST_KEY); err != nil {
				t.Fatalf("\t%s\tShould delete value at key (%s) successfully : %v", FAILED, TEST_KEY, err)
			}
			if v := k.Get(TEST_KEY); v != "" {
				t.Fatalf("\t%s\tShould delete value at key (%s) successfully. Value was not deleted.", FAILED, TEST_KEY)
			}
			t.Logf("\t%s\tShould delete value at key (%s) successfully.", SUCCEEDED, TEST_KEY)
		}

	}
}
