package queries

const (

	//login and register
	RegisterQuery       = `INSERT INTO register(id,name,password,email) values(?,?,?,?)`
	LoginQuery          = `select id, name, email from register where email=?`
	QueryForUniqueEmail = `select email from register`
	UpdateLoginState    = `update register set isLogin=true where id=?`
	LogOut              = `update register set isLogin=0 where email=? and password=?`
	//get ID of user
	GetIdOfUser = `select  id from register where name=? and password=?`
	//insert plan to DB , I think not in work
	InsertStudyPlan = `insert into plan (id,study) values(?,?)`
	//not in work
	GetPlan = `select study from plan where id=?`
	//getting user details for profile
	GetUserDetail = `select name, email from register where email=?`
	//save plan to DB
	SavePlanOfUser = `insert into plan(planSavedByUser,id,title,weekCount,CreatedAt) values(?,?,?,?,?)`
	//insert the 10 MCQs generated alongside the plan, tied to the new planID
	InsertMCQ = `insert into mcqTBL(mcqs, options, correctAns, planID) values(?,?,?,?)`
	//get id of user using email to save plan
	GetIdByEmailToSavePlan = `select id from register where email=?`
	//get id from email
	GetUserIDForProfile = `select id from register where email=?`
	//get each plan's title plus enough to work out completion: total weeks vs elapsed
	//time since creation, and total vs attempted MCQs. Used to show "my plan" list.
	GetPlanStatusForProfile = `select p.title as title, p.planID as planID, p.CreatedAt as createdAt, p.weekCount as weekCount,
		COUNT(m.mcqID) as totalMcq, COALESCE(SUM(m.isAttempted),0) as attemptedMcq
		from plan p left join mcqTBL m on m.planID = p.planID
		where p.id=?
		group by p.planID, p.title, p.CreatedAt, p.weekCount
		order by p.planID asc`
	// GetOwnPlanDescp  = `SELECT JSON_EXTRACT(planSavedByUser, '$.title','$.planID') AS title,planID FROM plan WHERE id=?`
	GetOwnPlanDescp = `select  planSavedByUser from plan where id=? and  title=? `
	//get week count to show the week count for particular plan
	// by sending title and id
	GetWeekCountTOshowAtDesp = `select weekCount from plan where id=? and title=?`
	//update the Now Time ---> this will click if user click on week box if a week is completed
	UpdateWeekCompleteTime = `Update plan set userClickNowTime=? where id=? and title=?`
	//get created and current time( when user click) to perform operation to calculate difference of 7 days
	GetBothTime = `select CreatedAt, userClickNowTime from plan where id=? and title=?`
	//update the completed week
	CompletedWeek = `Update plan set completedWeek=? where id=? and title=? `
	//get completeWeek count to update the week 1, week 2 .....count
	GetCompletedCount = `select completedWeek from plan where id=? and title=?`
	//get CreatedAt and total weekCount to calculate how many weeks are completed since plan creation (no DB write, computed on the fly)
	GetCreatedAtAndWeekCount = `select CreatedAt, weekCount from plan where id=? and title=?`
	//get planId to get MCQ for that particular plan
	GetPlanIDForParticularPlan = `select planID from plan where id=? and  title=?`
	//get one un-attempted MCQ against a planID, one at a time
	GetOneMcqTosolve = `select mcqID, mcqs, options from mcqTBL where planID=? and isAttempted=0 ORDER BY mcqID ASC limit 1`
	// is answer is correct or not [getting by planID]
	ValidatingAns = `select correctAns from mcqTBL where planID=?  and mcqID=?`
	//update  isCorrect  column from mcqTBL to true
	UpdateAnsStateAsTrue = `update mcqTBL set isCorrect=1  where planID=? and mcqID=?`
	//update  isCorrect  column from mcqTBL to false
	UpdateAnsStateAsFalse = `update mcqTBL set isCorrect=0 where planID=? and mcqID=?`
	//update  isAttempted =1 in mcqTBL
	UpdateAttemptedState = `update mcqTBL set isAttempted=1 where planID=? and mcqID=?`
	//store the option the user picked
	UpdateOptionSelected = `update mcqTBL set optionSelected=? where planID=? and mcqID=?`
	//get all 10 MCQs for a plan, to show the post-quiz review (correct/incorrect coloring)
	GetAllMcqsForPlan = `select mcqID, mcqs, options, correctAns, IFNULL(optionSelected,'') as optionSelected, IFNULL(isCorrect,0) as isCorrect from mcqTBL where planID=? order by mcqID asc`
	//get total vs attempted MCQ count for a plan, shown at the top of the quiz screen
	GetMcqCounts = `select count(*) as total, COALESCE(SUM(isAttempted),0) as attempted from mcqTBL where planID=?`

	// GetFirstMCQ = `select mcqs, from mcqTBL where planID=? ORDER BY mcqID ASC limit 1`
)
