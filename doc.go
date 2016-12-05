// Package zerconf is a pure Golang library that employs Multicast DNS-SD for
//
// * browsing and resolving services in your network
// * registering own services
//
// in the local network.
//
// It basically implements aspects of the standards
// [RFC 6762](https://tools.ietf.org/html/rfc6762) (mDNS) and
// [RFC 6763](https://tools.ietf.org/html/rfc6763) (DNS-SD).
// Though it does not support all requirements yet, the aim is to provide a complient solution in the long-term with the community.
//
// By now, it should be compatible to [Avahi](http://avahi.org/) (tested) and Apple's Bonjour (untested).
// ould work in the most office, home and private environments.
package zeroconf
