#include "scale_free.h"
#include <gsl/gsl_rng.h>
#include <pthread.h>
#include <time.h>

#define INTEGER_TYPE uint64_t

typedef struct parametric_module {
    struct parametric_module * previous;
    uint8_t kind;
    INTEGER_TYPE x;
    INTEGER_TYPE y;
} module;

typedef struct wrapper {
    module* m;
    gsl_rng* r;
} w;

module* pre_allocation;

void* rule( void* p) {
    #define M ((w *)p)->m
    #define R ((w *)p)->r
    // Memory of the particular block in memory, excuse the pun
    // struct parametric_module* end_of_block_pointer = m->previous;
    /* switch (m->kind) { */
    /*     case 'A': */
    // int r = ((m->y) - (m->x))/DIVISOR + m->x + 1; // Defining this here requires a memory call.
    module* elements = (module*) &pre_allocation[(CONNECTIONS + 1)*( ((M->y) - (M->x))/DIVISOR + M->x - 1)];
    /* module* elements = (module *)malloc((CONNECTIONS + 1)*sizeof(module)); */

    /* printf("%d",(CONNECTIONS + 1)*( ((M->y) - (M->x))/DIVISOR + M->x - 1)); */
    // ASSIGN the last element in the list
    // NOTES This MUST be the first element of the memory allocation block otherwise
    // WE CANNOT identify the start of the memory block to free later on...
    #define A_r elements[0]
    A_r.kind = 'A';
    A_r.x = M->x;
    A_r.y = ((M->y) - (M->x))/DIVISOR + M->x; // r-1
    A_r.previous = M->previous;

    M->previous = &elements[CONNECTIONS];
    M->x = A_r.y + 1;

    {
        INTEGER_TYPE last = A_r.y;
        double source = gsl_ran_beta(R, ALPHA, BETA);
        INTEGER_TYPE max = A_r.y + 1;
        for(int i =1; i < CONNECTIONS+1; i++) {
            elements[i].kind = 'L';
            elements[i].x = (INTEGER_TYPE)( source * last +1) % max + 1;
            elements[i].previous = &elements[i-1];
        }
    }

    #define check_M if (M->x != M->y) rule(p);

    if (A_r.x != A_r.y){
        w wrapper;
        wrapper.m = &A_r;
        if((A_r.y)-(A_r.x) > LIMIT ){
            pthread_t thread;
            wrapper.r = gsl_rng_alloc (gsl_rng_taus);
            pthread_create( &thread, NULL, rule, &wrapper);
            check_M
            pthread_join(thread,NULL);
            gsl_rng_free(wrapper.r);
        } else {
            wrapper.r = R;
            rule(&wrapper);
            check_M
        }
    } else check_M
    /*     default: */
    /*         break; */
    /* } */
    return 0;
}

int write_file(module* iv) {
    module* chain = iv;
    FILE *fptr;
    fptr = fopen("graph.adjlist", "w");
    int test = 0;
    do {
        switch (chain->kind) {
            case 'A':
                if( test++ == 0)
                    fprintf(fptr, "%d",chain->x);
                else
                    fprintf(fptr, "\n%d",chain->x);
                break;
            default:
                fprintf(fptr, " %d",chain->x);
                break;
        }
        chain = chain->previous;
    }while (chain != NULL);
    fclose(fptr);
    return 0;
}


int main(int argc, char *argv[]) {
    double total = 0.0;
    double* times = alloca(REPETITIONS*sizeof(double));
    for(int i = 0; i < REPETITIONS; i ++){


    gsl_rng *rand_src;
    rand_src = gsl_rng_alloc (gsl_rng_taus);
    INTEGER_TYPE max = MAX;
    module* iv = (module*)malloc(sizeof(module));
    iv->kind = 'A';
    iv->x = 1;
    iv->y = max;

    w wrapper;

    wrapper.m = iv;
    wrapper.r = rand_src;

    pre_allocation = (module *)malloc((CONNECTIONS + 1)*sizeof(module)*MAX);
    
    struct timespec start={0,0}, end={0,0};
    sleep(SECONDS_WAIT_BETWEEN_REPEATS);
    clock_gettime(CLOCK_MONOTONIC, &start);
    rule(&wrapper);
    clock_gettime(CLOCK_MONOTONIC, &end);
    times[i] = (end.tv_sec + 1.0e-9*end.tv_nsec) - (start.tv_sec + 1.0e-9*start.tv_nsec);
    total += times[i];
    printf("%.10fs\n",((end.tv_sec + 1.0e-9*end.tv_nsec) - (start.tv_sec + 1.0e-9*start.tv_nsec)));

    /* write_file(iv); */

    /* module* previous = iv; */
    /* do { */
    /*     iv = iv->previous; */
    /*     switch (iv->kind) { */
    /*         case 'A': */
    /*             free(previous); */
    /*             previous = iv; */
    /*     } */
    /* }while (previous->x != 1); */
    /* free(previous); */
    free(pre_allocation);
    gsl_rng_free(rand_src);
    }
    double average = total/(double)(REPETITIONS);
    double sum = 0;
    for (int i = 0; i < REPETITIONS; ++i)
        sum += pow(times[i] -average,2);
    double variance = sum/REPETITIONS;
    double std_deviation = sqrt(variance);
    printf("Average over %d : %.10fs : stdv: %.3f\%\n", (int)REPETITIONS, average,(std_deviation/average)*100);
    printf("%.2fmE/PE/s",(((double)(MAX*CONNECTIONS)/1000000)/PROCESSORS)/average);
    return 1;
}
