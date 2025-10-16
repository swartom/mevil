#ifndef SCALE_FREE_H_
#define SCALE_FREE_H_

#include <math.h>
#include <stdint.h>
#include <stdlib.h>
#include <gsl/gsl_randist.h>
#include <gsl/gsl_rng.h>
#include <stdio.h>
#include <unistd.h>

#define MAX 2500000
#define DIVISOR 2
#define CONNECTIONS 4
#define THREADS 12  // (uint32_t)sysconf(_SC_NPROCESSORS_ONLN) * 8
#define LIMIT MAX/THREADS
#define REPETITIONS 100
#define SECONDS_WAIT_BETWEEN_REPEATS 0

#define PROCESSORS 6

#define ALPHA 0.01
#define BETA 1.0

#endif // SCALE_FREE_H_2
