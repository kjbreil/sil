package silread

// Multi reads from sil files that contain multiple batches/tables points to a
// single SIL file and given a map of string to functions the string is the
// first part of the SIL view table. for example a OBJ_CHG (CHG is hard coded
// right now) would have a string OBJ_CHG
func Multi(filename string, tables map[string]interface{}) {

}
