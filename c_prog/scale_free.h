#ifndef SCALE_FREE_H_
#define SCALE_FREE_H_

#include <stdint.h>
#include <stdlib.h>
#include <gsl/gsl_randist.h>
#include <gsl/gsl_rng.h>
#include <stdio.h>
#include <unistd.h>

#define MAX 4000000
#define DIVISOR 2
#define CONNECTIONS 60
#define THREADS 64 // (uint32_t)sysconf(_SC_NPROCESSORS_ONLN) * 8
#define LIMIT MAX/THREADS

#define ALPHA 0.01
#define BETA 1.0

#endif // SCALE_FREE_H_2
