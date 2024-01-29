package settings

import "errors"

var (
	ErrCityNotValid                    = errors.New("the city specified is not valid")
	ErrControlServerPrivilegedPort     = errors.New("cannot use privileged port without running as root")
	ErrCountryNotValid                 = errors.New("the country specified is not valid")
	ErrFilepathMissing                 = errors.New("filepath is missing")
	ErrFirewallZeroPort                = errors.New("cannot have a zero port")
	ErrFirewallPublicOutboundSubnet    = errors.New("outbound subnet is public")
	ErrHostnameNotValid                = errors.New("the hostname specified is not valid")
	ErrISPNotValid                     = errors.New("the ISP specified is not valid")
	ErrMinRatioNotValid                = errors.New("minimum ratio is not valid")
	ErrMissingValue                    = errors.New("missing value")
	ErrNameNotValid                    = errors.New("the server name specified is not valid")
	ErrOpenVPNClientKeyMissing         = errors.New("client key is missing")
	ErrOpenVPNCustomPortNotAllowed     = errors.New("custom endpoint port is not allowed")
	ErrOpenVPNEncryptionPresetNotValid = errors.New("PIA encryption preset is not valid")
	ErrOpenVPNInterfaceNotValid        = errors.New("interface name is not valid")
	ErrOpenVPNKeyPassphraseIsEmpty     = errors.New("key passphrase is empty")
	ErrOpenVPNMSSFixIsTooHigh          = errors.New("mssfix option value is too high")
	ErrOpenVPNPasswordIsEmpty          = errors.New("password is empty")
	ErrOpenVPNTCPNotSupported          = errors.New("TCP protocol is not supported")
	ErrOpenVPNUserIsEmpty              = errors.New("user is empty")
	ErrOpenVPNVerbosityIsOutOfBounds   = errors.New("verbosity value is out of bounds")
	ErrOpenVPNVersionIsNotValid        = errors.New("version is not valid")
	ErrPortForwardingEnabled           = errors.New("port forwarding cannot be enabled")
	ErrPublicIPPeriodTooShort          = errors.New("public IP address check period is too short")
	ErrRegionNotValid                  = errors.New("the region specified is not valid")
	ErrServerAddressNotValid           = errors.New("server listening address is not valid")
	ErrSystemPGIDNotValid              = errors.New("process group id is not valid")
	ErrSystemPUIDNotValid              = errors.New("process user id is not valid")
	ErrSystemTimezoneNotValid          = errors.New("timezone is not valid")
	ErrUpdaterPeriodTooSmall           = errors.New("VPN server data updater period is too small")
	ErrVPNProviderNameNotValid         = errors.New("VPN provider name is not valid")
	ErrVPNTypeNotValid                 = errors.New("VPN type is not valid")
	ErrWireguardAllowedIPNotSet        = errors.New("allowed IP is not set")
	ErrWireguardAllowedIPsNotSet       = errors.New("allowed IPs is not set")
	ErrWireguardEndpointIPNotSet       = errors.New("endpoint IP is not set")
	ErrWireguardEndpointPortNotAllowed = errors.New("endpoint port is not allowed")
	ErrWireguardEndpointPortNotSet     = errors.New("endpoint port is not set")
	ErrWireguardEndpointPortSet        = errors.New("endpoint port is set")
	ErrWireguardInterfaceAddressNotSet = errors.New("interface address is not set")
	ErrWireguardInterfaceAddressIPv6   = errors.New("interface address is IPv6 but IPv6 is not supported")
	ErrWireguardInterfaceNotValid      = errors.New("interface name is not valid")
	ErrWireguardPreSharedKeyNotSet     = errors.New("pre-shared key is not set")
	ErrWireguardPrivateKeyNotSet       = errors.New("private key is not set")
	ErrWireguardPublicKeyNotSet        = errors.New("public key is not set")
	ErrWireguardPublicKeyNotValid      = errors.New("public key is not valid")
	ErrWireguardImplementationNotValid = errors.New("implementation is not valid")
)
