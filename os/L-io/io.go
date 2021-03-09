package L_io

import "io"

// os.File 同时实现了 io.Reader 和 io.Writer
// strings.Reader 实现了 io.Reader
// bufio.Reader/Writer 分别实现了 io.Reader 和 io.Writer
// bytes.Buffer 同时实现了 io.Reader 和 io.Writer
// bytes.Reader 实现了 io.Reader
// compress/gzip.Reader/Writer 分别实现了 io.Reader 和 io.Writer
// crypto/cipher.StreamReader/StreamWriter 分别实现了 io.Reader 和 io.Writer
// crypto/tls.Conn 同时实现了 io.Reader 和 io.Writer
// encoding/csv.Reader/Writer 分别实现了 io.Reader 和 io.Writer
// mime/multipart.Part 实现了 io.Reader
// net/conn 分别实现了 io.Reader 和 io.Writer(Conn接口定义了Read/Write)

type IO struct {}

// Read 将 len(p) 个字节读取到 p 中。 数据从 IO -> p
// 它返回读取的字节数 n (0 <= n <= len(p)) 和 任何遇到的错误。
// 即使 Read 返回的 n < len(p)，它也会在调用过程中占用 len(p) 个字节作为暂存空间。
// 若可读取的数据不到 len(p) 个字节，Read 会返回可用数据，而不是等待更多数据。
//
// 当 Read 在成功读取 n > 0 个字节后遇到一个错误或 EOF (end-of-file) 它会返回读取的字节数。
// 它可能会同时在本次的调用中返回一个non-nil错误, 或在下一次的调用中返回这个错误（且 n 为 0）。
// 一般情况下, Reader会返回一个非0字节数n, 若 n = len(p) 个字节从输入源的结尾处由 Read 返回，
// Read可能返回 err == EOF 或者 err == nil。
// 并且之后的 Read() 都应该返回 (n:0, err:EOF)。
//
// 调用者在考虑错误之前应当首先处理返回的数据。
// 这样做可以正确地处理在读取一些字节后产生的 I/O 错误，同时允许EOF的出现。
func (i *IO) Read(p []byte) (n int, err error) {
	return 0, err
}

// Write 将 len(p) 个字节从 p 中写入到基本数据流中。 数据从 p -> IO
// 它返回从 p 中被写入的字节数 n（0 <= n <= len(p)）以及任何遇到的引起写入提前停止的错误。
// 若 Write 返回的 n < len(p)，它就必须返回一个 非nil 的错误。
func (i *IO) Write(p []byte) (n int, err error) {
	return 0, err
}

// ReadAt 从基本输入源的偏移量 off 处开始，将 len(p) 个字节读取到 p 中。
// 它返回读取的字节数 n（0 <= n <= len(p)）以及任何遇到的错误。
//
// 当 ReadAt 返回的 n < len(p) 时，它就会返回一个 非nil 的错误来解释 为什么没有返回更多的字节。
// 在这一点上，ReadAt 比 Read 更严格。
//
// 即使 ReadAt 返回的 n < len(p)，它也会在调用过程中使用 p 的全部作为暂存空间。
// 若可读取的数据不到 len(p) 字节，ReadAt 就会阻塞, 直到所有数据都可用或一个错误发生。
// 在这一点上 ReadAt 不同于 Read。
//
// 若 n = len(p) 个字节从输入源的结尾处由 ReadAt 返回，Read可能返回 err == EOF 或者 err == nil
//
// 若 ReadAt 携带一个偏移量从输入源读取，ReadAt 应当既不影响偏移量也不被它所影响。
//
// 可对相同的输入源并行执行 ReadAt 调用。
func (i *IO) ReadAt(p []byte, off int64) (n int, err error) {
	return 0, err
}

// WriteAt 从 p 中将 len(p) 个字节写入到偏移量 off 处的基本数据流中。
// 它返回从 p 中被写入的字节数 n（0 <= n <= len(p)）以及任何遇到的引起写入提前停止的错误。
// 若 WriteAt 返回的 n < len(p)，它就必须返回一个 非nil 的错误。
//
// 若 WriteAt 携带一个偏移量写入到目标中，WriteAt 应当既不影响偏移量也不被它所影响。
//
// 若被写区域没有重叠，可对相同的目标并行执行 WriteAt 调用。
func (i *IO) WriteAt(p []byte, off int64) (n int, err error) {
	return 0, err
}

// ReadFrom 从 r 中读取数据，直到 EOF 或发生错误。其返回值 n 为读取的字节数。
// 除 io.EOF 之外，在读取过程中遇到的任何错误也将被返回。
//
// 如果 ReaderFrom 可用，Copy 函数就会使用它。
func (i *IO) ReadFrom(r io.Reader) (n int64, err error) {
	return 0, err
}

// WriteTo 将数据写入 w 中，直到没有数据可写或发生错误。
// 其返回值 n 为写入的字节数。
// 在写入过程中遇到的任何错误也将被返回。
//
// 如果 WriterTo 可用，Copy 函数就会使用它。
func (i *IO) WriteTo(w io.Writer) (n int64, err error) {
	return 0, err
}

// Seek 设置下一次 Read 或 Write 的偏移量为 offset
// 它的解释取决于 whence：
// 0 表示相对于文件的起始处，
// 1 表示相对于当前的偏移，
// 2 表示相对于其结尾处。
// Seek 返回新的偏移量和一个错误，如果有的话。
func (i *IO) Seek(offset int64, whence int) (ret int64, err error) {
	return 0, err
}

// 用于关闭数据流。
// 文件 (os.File)、归档（压缩包）、数据库连接、Socket 等需要手动关闭的资源都实现了 Closer 接口。
// 实际编程中，经常将 Close 方法的调用放在 defer 语句中。
func (i *IO) Close() error {
	return nil
}

// ReadFrom 函数将 io.Reader 作为参数，
// 也就是说，ReadFrom 可以从任意的地方读取数据，只要来源实现了 io.Reader 接口。
// 比如，我们可以从标准输入、文件、字符串等读取数据
func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}