let fibonacci = fn(x) {
	if (x == 0) {1} else {
		if (x == 1) {1} else {
			fibonacci(x - 1) + fibonacci(x-2)
		}
	}
};

puts(fibonacci(15));
