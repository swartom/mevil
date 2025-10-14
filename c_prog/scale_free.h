#ifndef SCALE_FREE_H_
#define SCALE_FREE_H_

#include <stdint.h>
#include <stdlib.h>
#include <gsl/gsl_randist.h>
#include <gsl/gsl_rng.h>
#include <stdio.h>

#define MAX 1000000
#define DIVISOR 2
#define CONNECTIONS 60
#define THREADS 1024
#define LIMIT MAX/THREADS

#define ALPHA 0.5
#define BETA 1.0

#endif // SCALE_FREE_H_
