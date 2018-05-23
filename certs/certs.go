package certs;

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
)

var (
	caPem = []byte(`
-----BEGIN CERTIFICATE-----
MIIC9zCCAd+gAwIBAgIJALGtqdMzpDemMA0GCSqGSIb3DQEBCwUAMBIxEDAOBgNV
BAMMB2t1YmUtY2EwHhcNMTYxMDE5MTU1NDI0WhcNNDQwMzA2MTU1NDI0WjASMRAw
DgYDVQQDDAdrdWJlLWNhMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA
pa4Wu/WkpJNRr8pMVE6jjwzNUOx5mIyoDr8WILSxVQcEeyVPPmAqbmYXtVZO11p9
jTzoEqF7Kgts3HVYGCk5abqbE14a8Ru/DmV5avU2hJ/NvSjtNi/O+V6SzCbg5yR9
lBR53uADDlzuJEQT9RHq7A5KitFkx4vUcXnjOQCbDogWFoYuOgNEwJPy0Raz3NJc
ViVfDqSJ0QHg02kCOMxcGFNRQ9F5aoW7QXZXZXD0tn3wLRlu4+GYyqt8fw5iNdLJ
t79yKp8I+vMTmMPz4YKUO+eCl5EY10Qs7wvoG/8QNbjH01BRN3L8iDT2WfxdvjTu
1RjPxFL92i+B7HZO7jGLfQIDAQABo1AwTjAdBgNVHQ4EFgQUZTrg+Xt87tkxDhlB
gKk9FdTOW3IwHwYDVR0jBBgwFoAUZTrg+Xt87tkxDhlBgKk9FdTOW3IwDAYDVR0T
BAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEApB7JFVrZpGSOXNO3W7SlN6OCPXv9
C7rIBc8rwOrzi2mZWcBmWheQrqBo8xHif2rlFNVQxtq3JcQ8kfg/m1fHeQ/Ygzel
Z+U1OqozynDySBZdNn9i+kXXgAUCqDPp3hEQWe0os/RRpIwo9yOloBxdiX6S0NIf
VB8n8kAynFPkH7pYrGrL1HQgDFCSfa4tUJ3+9sppnCu0pNtq5AdhYx9xFb2sn+8G
xGbtCkhVk2VQ+BiCWnjYXJ6ZMzabP7wiOFDP9Pvr2ik22PRItsW/TLfHFXM1jDmc
I1rs/VUGKzcJGVIWbHrgjP68CTStGAvKgbsTqw7aLXTSqtPw88N9XVSyRg==
-----END CERTIFICATE-----
`)
	keyPem = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAqoDhGW3KQWB+fJFud4j84DQqpIbtEat7MAaoEAEkLakVAfis
RSvxUQVI+GjA4B11NAdHM3eaVVSvz41ydHbTm9989GrTC5O7kW8eNkHMf4boWbns
9R58JfcjnQLUpB2m7Oi50nUeyPZSrGZHtEVCbgFjXADyZtn06DN1ZVlQMJRaA4E5
66yhXqZp1tlvevZJMSfIO7MexVTHwFbxolJ8ju/NHeFdu1G9s0T8SN7imkjPjsGc
h8Oo7TnMjE8UJyAEIfJjybDFntwTurWyvW8DJemvtFUEYWjAfjll0qrPURUmEuhk
CRIynxh+1SFMG8KRxcnX4rOS70UybqA2ZIHOhQIDAQABAoIBADkacskucQeWRqZ3
mpSkJ3T7Y7C4k5tQYDCEejLp/vDf6O5BF4kPH4HwEDFJ/BbTJtam/VyqWODtPPh5
OfTxewuAPPwq7sW405/wpCCtxsyLJTQvxcGOVSvt6lqCgS7501cS1nE60nWhsayV
kLw6WfY3tswwcP6rTo+Z9F5eHDI4YQ3yKkaIy3FKGfGxo35YStnj3EQkezIktMSR
KDJVhBzbz4ydai0enxcMQTs7IRPDXSuBfAmYuEXr0k7FofkWGDRa3nYIoA+99hQ+
s/Tid004dj/+lqSVill26NNXe+OdUAAsHHIJ9cUWeB5X+0pUQypHbBL6hBWVycnu
38kt430CgYEA3D3eEr0XGCVyteN9NyH6SDfwjb7YWzeOi6lS3cYwbDhh0uxnlgOs
TI4P6tQeT2Tl1PtV6fyKBCvG7ywv47gIvLeS7T+0S8qqO+lHWzLXaVes1zJ1jlxK
sO3tKxu/9xZ9ep2pjZMJurO4/86hdfwhQV2/q+3HhY7Mm8jwN9xpl6MCgYEAxi+w
qBza3RIWlp6bbv2mmiKZ0T+x0kcDffAzdQKjJi3lmYWjpHzXfLgjiOsqo27/qm0u
MKLMu3TyhJwtaOxO9OdDhQmqG5uOlLFb9UU2nhEAcnh2YMRrHCj695C9S97PkLVJ
s2kp4uVFMVxpKsLgVlwu0GH3zugPQH4v/eC6g7cCgYB7dDKHTncjkdo7GsmVnfYt
hS3SRqgAeaPtpXxN1EpQX0p2cQ2fiW+LehZyC3TyDBzIxhnijyzOKbvZVWuCuiYr
ors5QfxOf8vsyVa2SEl3Qy4fcqlyo0k65CONhoCUgIbVtIrWURWjEhshSTI4cJwz
h9lpBmBQ/Tq0GG6O4X5PAQKBgQCY8Ah1Uv8ahnDj/rWX6yn73COzNGH3EVICh7BN
5aEdP2HlHRnxP13TIw5ZBJE82dV2IRb59UfkiRT1fMgWJfWwTB7wtUqOT3ayDEQY
fDbvt9MOgyNm/WxiqMUy8oEB4Ylv9FZRmx/1tlO1CckmdIhGXJDLwi5HfxD2Bern
Edsc9QKBgBI1dOaD/6UH9U0Q/tmHIdYBtB6ZM0nn/zLl91PBtUCLI/V83S5O66wf
NPzzPY7LzRc61z1pOx718NgnYSH7HLjjz5xtClctLk8KPZcWvSE7QoFa1vDWLOGA
yErXmrD43yQ0azVLHAWJdxIUjgUjRDv/Xn27JRGzizP5W164zAiS
-----END RSA PRIVATE KEY-----
`)
	certPem = []byte(`
-----BEGIN CERTIFICATE-----
MIIC9jCCAd6gAwIBAgIJAJy0KUReH+YSMA0GCSqGSIb3DQEBCwUAMBIxEDAOBgNV
BAMMB2t1YmUtY2EwHhcNMTcwMTEyMjIxMjAwWhcNMjAxMDA4MjIxMjAwWjAUMRIw
EAYDVQQDDAlsb2NhbGhvc3QwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIB
AQCqgOEZbcpBYH58kW53iPzgNCqkhu0Rq3swBqgQASQtqRUB+KxFK/FRBUj4aMDg
HXU0B0czd5pVVK/PjXJ0dtOb33z0atMLk7uRbx42Qcx/huhZuez1Hnwl9yOdAtSk
Habs6LnSdR7I9lKsZke0RUJuAWNcAPJm2fToM3VlWVAwlFoDgTnrrKFepmnW2W96
9kkxJ8g7sx7FVMfAVvGiUnyO780d4V27Ub2zRPxI3uKaSM+OwZyHw6jtOcyMTxQn
IAQh8mPJsMWe3BO6tbK9bwMl6a+0VQRhaMB+OWXSqs9RFSYS6GQJEjKfGH7VIUwb
wpHFydfis5LvRTJuoDZkgc6FAgMBAAGjTTBLMAkGA1UdEwQCMAAwCwYDVR0PBAQD
AgXgMDEGA1UdEQQqMCiCCWxvY2FsaG9zdIIVbG9jYWxob3N0LmxvY2FsZG9tYWlu
hwR/AAABMA0GCSqGSIb3DQEBCwUAA4IBAQB0XD8uRAxtch9NKqgQ1QD9jNygEDpt
POwvCj4rs39mSf4DIKFv9y+lTzTw/0P7mSUgoc5klekmStY8ql3HtmXo03MrK964
XTuFdpCePYKHQn5ylYzqRCDaVhjHljI6CUgc6UKpBE+JUqr6fznztYQU5Qtgacml
RIk8sWzS/S5fF65kaRRGgL88p28zZY760smV2VrtFHCN1+atFuW8QnuvtXTsQUFf
KpU8xcb7PfK7dbNa9vkxGa0LSg/qxddAdv0QrN7pUMipsSWQz5qh1z0TtIXjRnPo
hPLV3L4dhWqepPVP/el6kkZftmY2lJrRTcI1WNIILZRYSSvM0+ovAmnM
-----END CERTIFICATE-----
`)
)

// NewTLSConfig returns a TLS config that includes a certificate
// Use for server TLS config or when using a client certificate
// If any path is empty, a default PEM will be used
func NewTLSConfig(certPath, keyPath, caPath string) (*tls.Config, error) {
	var cert tls.Certificate
	var err error
	if certPath != "" && keyPath != "" {
		cert, err = tls.LoadX509KeyPair(certPath, keyPath)
	} else {
		cert, err = tls.X509KeyPair(certPem, keyPem)
	}
        if err != nil {
                return nil, err
        }

        roots, err := loadRoots(caPath)
        if err != nil {
                return nil, err
        }

        return &tls.Config{Certificates: []tls.Certificate{cert}, RootCAs: roots}, nil
}

func NewServerTLSConfig(certPath, keyPath, caPath string) (*tls.Config, error) {
	c, err := NewTLSConfig(certPath, keyPath, caPath)
	if err != nil {
		return nil, err
	}
	c.ClientAuth = tls.RequireAndVerifyClientCert
	c.ClientCAs = c.RootCAs
	return c, nil
}

func loadRoots(caPath string) (*x509.CertPool, error) {
        roots := x509.NewCertPool()

	pem := caPem
        if caPath != "" {
		pemtxt, err := ioutil.ReadFile(caPath)
		if err != nil {
			return nil, err
		}
		pem = pemtxt
        }
        ok := roots.AppendCertsFromPEM(pem)
        if !ok {
                return nil, fmt.Errorf("Could not append cert to roots")
        }
        return roots, nil
}
