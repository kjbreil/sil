package loc

// TaskTab is the TASK_TAB definition
type TaskTab struct {
	UPCCode                 string  `sil:"F01,zeropad"`
	TargetIdentifier        string  `sil:"F1000"`
	RecordStatus            int     `sil:"F1001" default:"1"`
	TerminalStore           string  `sil:"F1056"`
	UserRestrictionLevel    *int    `sil:"F1142"`
	CustomerId              *string `sil:"F1148"`
	OperatorResponsible     *int    `sil:"F1168"`
	BuyingFormat            *string `sil:"F1184"`
	UserOrderNumber         *int    `sil:"F1185"`
	ReferenceNumber         *string `sil:"F1245"`
	EventRegularType        *string `sil:"F1675"`
	EventRegularDiscount    *string `sil:"F1676"`
	EventPromoType          *string `sil:"F1677"`
	EventPromoDiscount      *string `sil:"F1678"`
	EventRegularDays        *int    `sil:"F1679"`
	EventPromoDate          *string `sil:"F1680"`
	TaskID                  *string `sil:"F1822"`
	TaskType                *string `sil:"F1823"`
	TaskDescription         *string `sil:"F1824"`
	TaskURL                 *string `sil:"F1825"`
	TaskDueDate             *string `sil:"F1826"`
	TaskPrioriry            *int    `sil:"F1827"`
	TaskRecurseFormulas     *string `sil:"F1828"`
	TaskEnabled             *string `sil:"F1829"`
	LastChangeDate          string  `sil:"F253" default:"NOW"`
	VendorId                *string `sil:"F27"`
	OperatorLastTouch       *int    `sil:"F2754"`
	TaskComments            *string `sil:"F2755"`
	TaskLastComment         *string `sil:"F2756"`
	TaskTimeSpent           *string `sil:"F2757"`
	TaskCategory            *string `sil:"F2758"`
	TaskCategory2           *string `sil:"F2759"`
	TaskDateCreation        *string `sil:"F2760"`
	TaskContactHost         *int    `sil:"F2761"`
	TaskContactRemote       *int    `sil:"F2762"`
	TaskReferenceNote       *string `sil:"F2763"`
	TaskReferenceDate       *string `sil:"F2764"`
	TaskExecutionDateOffset *int    `sil:"F2791"`
}
