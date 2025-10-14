#include "scale_free.h"
#include <gsl/gsl_rng.h>
#include <pthread.h>
#include <time.h>

typedef struct parametric_module {
    struct parametric_module * previous;
    uint8_t kind;
    uint32_t x;
    uint32_t y;
} module;
gsl_rng *rand_src;

void* rule( void* ptr) {
    module* m = (module *) ptr;
    // Memory of the particular block in memory, excuse the pun
    // struct parametric_module* end_of_block_pointer = m->previous;
    switch (m->kind) {
        case 'A':

            // int r = ((m->y) - (m->x))/DIVISOR + m->x + 1; // Defining this here requires a memory call.
            module* elements = (module *)malloc((CONNECTIONS + 1)*sizeof(module));
            // Assign the last element in the list
            // NOTES This MUST be the first element of the memory allocation block otherwise
            // WE CANNOT identify the start of the memory block to free later on...
            module *A_r = &elements[0];
            A_r->kind = 'A';
            A_r->x = m->x;
            // This value is = r - 1
            A_r->y = ((m->y) - (m->x))/DIVISOR + m->x;
            A_r->previous = m->previous;
            module *repeater_pointer = A_r;
            /* if (A_r->y  < CONNECTIONS) { */
            /*     // Case for there is fewer connections than there are differences */
            /*     for(int i =0; i < CONNECTIONS; i++) { */
            /*         // Each element value assigned. */
            /*         elements[i].kind = 'L'; */
            /*         elements[i].x = m->x -i - 1; */
            /*         elements[i].previous = repeater_pointer; */
            /*         repeater_pointer = &elements[i]; */
            /*         // Filter Through the repeater pointer */
            /*     } */
            /* } else { */
                // There should be *some* form of stochastic process selection...
            for(int i =1; i < CONNECTIONS+1; i++) {
                // Each element value assigned.
                elements[i].kind = 'L';
                elements[i].x = (uint32_t)(gsl_ran_beta(rand_src, ALPHA, BETA) * (double)(A_r->y + 1)) + 1;
                elements[i].previous = repeater_pointer;
                repeater_pointer = &elements[i];
            }

            /* } */
            // Link up to list!
            m->previous = repeater_pointer;
            m->x = A_r->y + 1;

            if (A_r->x != A_r->y){
                if((A_r->y)-(A_r->x) >= LIMIT ){
                    pthread_t thread;
                    pthread_create( &thread, NULL, rule, A_r);
                    if (m->x != m->y){
                        rule(m);
                    }
                    pthread_join(thread,NULL);
                } else {
                    rule(A_r);
                    if (m->x != m->y){
                        rule(m);
                    }
                }
            } else if (m->x != m->y){
                rule(m);
            }
        default:
            break;
    }
    return 0;
}

int write_file(module* iv) {
    module* chain = iv;
    FILE *fptr;
    fptr = fopen("graph.adjlist", "w");
    do {
        switch (chain->kind) {
            case 'A':
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
    for(int i =0; i < 10; i++) {
    rand_src = gsl_rng_alloc (gsl_rng_taus);
    uint32_t max = MAX;
    module* iv = (module*)malloc(sizeof(module));
    iv->kind = 'A';
    iv->x = 1;
    iv->y = max;

    struct timespec start={0,0}, end={0,0};
    clock_gettime(CLOCK_MONOTONIC, &start);
    rule(iv);
    clock_gettime(CLOCK_MONOTONIC, &end);
    printf("%.5f\n",((double)end.tv_sec + 1.0e-9*end.tv_nsec) - (
               (double)start.tv_sec + 1.0e-9*start.tv_nsec));

    /* write_file(iv); */

    module* previous = iv;
    do {
        iv = iv->previous;
        switch (iv->kind) {
            case 'A':
                free(previous);
                previous = iv;
        }
    }while (previous->x != 1);
    free(previous);
    gsl_rng_free(rand_src);
    }
    return 1;
}
