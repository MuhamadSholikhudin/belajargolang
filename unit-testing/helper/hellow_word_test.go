

package helper
import "testing"
func TestHelloWorld(t testing.T) {
result := HelloWorld("Eko")
if result != "Hello Eko" { // error
panic "Result is not "Hello Eko")
}
}