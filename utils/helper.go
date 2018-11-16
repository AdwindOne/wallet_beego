/**
 * Created by Adwind.
 * User: liuyunlong
 * Date: 11/15/18
 * Time: 17:13
 */
package utils

import "encoding/json"

func DecodeByte(data []byte) interface{}  {
	var v interface{}
	json.Unmarshal(data,&v)
	return v
}
