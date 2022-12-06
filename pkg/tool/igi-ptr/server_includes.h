#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#include <sys/param.h>
#include <sys/socket.h>
#include <fcntl.h>
#include <sys/stat.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <netdb.h>
#include <string.h>
#include <unistd.h>
#include <signal.h>
#include <sys/time.h>
#include <pthread.h>

#include "common.h"

struct filter_item                                                                                                       
{
    char src_ip_str[NI_MAXHOST];
    char dst_ip_str[NI_MAXHOST];
    /* sockaddr型とかAFを一緒に格納できた方が良い */
    char src_ip[NI_MAXHOST];
    char dst_ip[NI_MAXHOST];
    int control_sock;
 
    int listen_port;
    int listen_sock;
 
    struct pkt_rcd_t record[MaxProbeNum];
    /* last active time */
    double pre_time;
    /* the number of packets we have filtered out for this filter */
    int count;
    /* the number of packets client should send to us */
    int probe_num;
 
    /* number of consecutive no data phase, 3 means dead */
    int nodata_count;
    /* working state */
    int dead;
    pthread_t thread;
 
    struct filter_item *next;
};

void Usage();
double get_time();
void phase_finish();
void get_packets(struct filter_item *ptr);
void update_filter_list(
     int newsd,
     int probe_num,
     char *src_ip_str,
     char *dst_ip_str,
     int *listening_port);
void check_client();
void main_server();

