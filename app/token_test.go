package main

import (
	"log"
	"testing"
)

func TestValidateJWT(t *testing.T) {

	t.Log("when trying to validate jwt")

	hmacSampleSecret := "jdnfksdmfksd"

	// step 1 - generate jwt
	token, err := NewAuthToken("123", hmacSampleSecret)
	if err != nil {
		t.Errorf("cant generate jwt %v", err)
	}
	log.Printf("%v", token)

	// step 2 - validate jwt
	_, err = ValidateJWT(token.Token, hmacSampleSecret)

	// check result
	var expected error
	expected = nil
	result := err
	t.Logf("expecting to receive %v", expected)
	if result != expected{
		t.Errorf("expected %v but received %v", expected, result)
	}
}

func TestVerifyClaims(t *testing.T) {

	t.Log("when trying to validate jwt")

	hmacSampleSecret := "jdnfksdmfksd"

	// step 1 - generate jwt
	token, err := NewAuthToken("123", hmacSampleSecret)
	if err != nil {
		t.Errorf("cant generate jwt %v", err)
	}

	// step 2 - validate jwt
	err = VerifyClaims(token.Token, hmacSampleSecret)

	// check result
	var expected error
	expected = nil
	result := err
	t.Logf("expecting to receive %v", expected)
	if result != expected{
		t.Errorf("expected %v but received %v", expected, result)
	}
}