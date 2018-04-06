#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "kflow.h"

int main(int argc, char **argv) {
    char *err;
    int r;

    kflowConfig cfg = {
        .URL = "http://127.0.0.1:8999/chf",
        .API = {
            .email = "test@example.com",
            .token = "token",
            .URL   = "http://127.0.0.1:8999/api/internal",
        },
        .metrics = {
            .interval = 1,
            .URL      = "http://127.0.0.1:8999/tsdb",
        },
        .device_id = 1,
        .verbose   = 1,
        .program   = "demo",
        .version   = "0.0.1",
    };

    kflowDevice device;
    kflowCustom *customs;
    uint32_t numCustoms;

    if ((r = kflowInit(&cfg, &device)) != 0) {
        printf("error initializing libkflow: %d\n", r);
        goto error;
    };

    customs    = device.customs;
    numCustoms = device.num_customs;

    char *url = "http://foo.com";

    for (uint32_t i = 0; i < numCustoms; i++) {
        if (!strcmp(customs[i].name, KFLOWCUSTOM_HTTP_URL)) {
            customs[i].value.str = url;
        } else if (!strcmp(customs[i].name, KFLOWCUSTOM_HTTP_STATUS)) {
            customs[i].value.u32 = 200;
        } else {
            free(customs[i].name);
            memmove(&customs[i], &customs[i+1], sizeof(kflowCustom)*(numCustoms-i-1));
            numCustoms--; i--;
        }
    }

    kflow flow = {
        .deviceId    = cfg.device_id,
        .ipv4SrcAddr = 167772161,
        .ipv4DstAddr = 167772162,
        .srcAs       = 1234,
        .inPkts      = 20,
        .inBytes     = 40,
        .srcEthMac   = 1250999896491,
        .dstEthMac   = 226426397786884,
        .customs     = customs,
        .numCustoms  = numCustoms,
    };

    if ((r = kflowSend(&flow)) != 0) {
        printf("error sending flow: %d\n", r);
        goto error;
    }

    if ((r = kflowStop(10*1000)) != 0) {
        printf("error stopping libkflow: %d\n", r);
        goto error;
    }

    return 0;

  error:

    while ((err = kflowError()) != NULL) {
        printf("libkflow error: %s\n", err);
        free(err);
    }

    exit(1);
}
