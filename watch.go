package glob

// WatchFunc function
type WatchFunc func(path string)

// Watch function
func Watch(glob string, fn WatchFunc) error {
	return nil
}
