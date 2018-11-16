/**
 * Created by Adwind.
 * User: liuyunlong
 * Date: 11/13/18
 * Time: 16:38
 */
package utils

type ReUtil struct {
	Code int
	Message string
	Data interface{}
}
func NewReutil() *ReUtil {
	re := new(ReUtil)
	return re
}