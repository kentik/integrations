#include <arpa/inet.h>
#include <stdint.h>

// returns a string version of the address. Uses memory passed in as a buffer
char* fast_intoa (unsigned int addr, char* buf, int bufLen) {
    char * cp, * retStr;
    uint8_t byte;
    int n;

    cp = &buf[bufLen];
    *--cp = '\0';

    n = 4;
    do
        {
            byte = addr & 0xff;
            *--cp = byte % 10 + '0';
            byte /= 10;
            if (byte > 0)
                {
                    *--cp = byte % 10 + '0';
                    byte /= 10;
                    if (byte > 0)
                        *--cp = byte + '0';
                }
            *--cp = '.';
            addr >>= 8;
        }
    while (--n > 0);

    /* Convert the string to lowercase */
    retStr = (char*)(cp+1);

    return(retStr);
}

// Returns the packed representaion of this string.
u_int32_t pack_ipv4_address(const char* ipv4)
{

    u_int32_t address;

    /* Converts the Internet host address IPv4 from numbers-and-dots notation
     * into binary data in network byte order. */
    address = inet_addr (ipv4);

    /* Converts the address from network byte order to host byte order */
    return ntohl (address);
}
