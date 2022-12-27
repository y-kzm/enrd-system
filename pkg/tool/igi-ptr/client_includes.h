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

void Usage_client();
double get_rcd_time(struct pkt_rcd_t record);
void dump_trace();
void quit();
double get_time_client();
int get_delay_num(double gap);
void dump_bandwidth();
int init_sockets(struct sockaddr *dst_ip);
int get_host_info(char *string, char *name, uint32_t *ip, char *ip_str);
void get_item(int control_sock, char *str);
void init_connection();
void send_packets(int probe_num, int packet_size, int delay_num, double *sent_times);
double get_dst_sum(struct pkt_rcd_t *rcv_record, int count, int *gap_count);
int get_dst_gaps(struct pkt_rcd_t *rcv_record);
double get_bottleneck_bw(struct pkt_rcd_t *rcv_record, int count);
double get_competing_bw(struct pkt_rcd_t *rcv_record, double avg_src_gap, int count, double b_bw);
double get_src_sum(double *times);
void get_bandwidth();
void one_phase_probing();
void n_phase_probing(int n);
int gap_comp(double dst_gap, double src_gap);
void fast_probing();
void init_connection();
void send_packets(int probe_num, int packet_size, int delay_num, double *sent_times);
double get_dst_sum(struct pkt_rcd_t *rcv_record, int count, int *gap_count);
double main_client();
//double main_client(int phase_num, int probe_num, int packet_size, char src_addr, char dst_addr);
