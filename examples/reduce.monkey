let reduce = fn(arr, func, start) {
  let acc = fn(arr, func, buf) {
    if (len(arr) == 0) {buf} else {
      let buf = func(buf, arr[0]);
      return acc(rest(arr), func, buf);
    }
  }
  return acc(arr, func, start);
}

let add = fn(a, b) {a + b}

puts(reduce([1, 2, 3], add, 0))
