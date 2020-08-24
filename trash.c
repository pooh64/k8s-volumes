#include <stdlib.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>
#include <string.h>
#include <time.h>

void inject(char *mem)
{
	for (int i = 0; i < 32; ++i)
		mem[rand() % 4096] = rand();
}

int main(int argc, char **argv)
{
	if (argc != 2)
		return 1;

	long sz = atol(argv[1]);
	char *mem = malloc(4096);
	memset(mem, 32, 4096);
	int fd = open("trash", O_WRONLY | O_APPEND);
	if (fd < 0)
		return 1;

	srand(time(NULL));

	for (long i = 0; i < sz / 4096L; ++i) {
		inject(mem);
		write(fd, mem, 4096);
	}

	return 0;
}

