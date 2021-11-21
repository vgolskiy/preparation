package support

import (
	"fmt"
	"os"
)

//type s_ph struct {
//	pthread_t		ph_id;
//	pthread_t		checker_id;
//	int				number;
//	int				eat_qty;
//	t_bool			fork_left;
//	t_bool			fork_right;
//	t_bool			eating;
//	uint64_t		time_limit_life;
//	struct s_st		*st;
//}
//
//typedef struct	s_st
//{
//int				qty;
//int				eat_max;
//t_bool			stop;
//uint64_t		time_die;
//uint64_t		time_eat;
//uint64_t		time_sleep;
//uint64_t		time_start;
//pthread_mutex_t	mutex_status;
//pthread_mutex_t	mutex_print;
//pthread_mutex_t	*mutex_forks;
//struct s_ph		*ph;
//}				t_st;

var err = map[int]string{
	1: "Wrong number of arguments",
}

func error(err_code int) {
	fmt.Errorf(err[err_code])
	os.Exit(1)
}
