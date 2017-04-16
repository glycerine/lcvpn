package main

// encnone implements just copy "enctrypyion"-"decryption" (no encryption at all!)
type encnone struct {
}

func newEncNone(key string) (PacketEncrypter, error) {

	a := encnone{}
	return &a, nil

}

func (a *encnone) CheckSize(size int) bool {
	return true
}

func (a *encnone) AdjustInputSize(size int) int {
	return size
}

func (a *encnone) Encrypt(input []byte, output []byte, iv []byte) int {
	copy(output, input)
	return len(input)
}

func (a *encnone) Decrypt(input []byte, output []byte) int {
	copy(output, input)
	return len(input)
}

func (a *encnone) OutputAdd() int {
	return 0
}

func (a *encnone) IVLen() int {
	return 0
}

func init() {
	if registredEncrypters == nil {
		registredEncrypters = make(map[string]newEncrypterFunc)
	}
	registredEncrypters["none"] = newEncNone
}
