package main


func init() {
  // Daily Users
  RegisterReport("users",
    "SELECT MONTH(date_joined) date, COUNT(id) total FROM auth_user WHERE 1=1 AND YEAR(date_joined) = 2013 GROUP BY 1",
    []string{"%d,","%d"},
    "Get monthly 2013 users",
  )

  // Daily sent reservations
  RegisterReport("funnel_1",
    "SELECT WEEKOFYEAR(start_date) date, COUNT(id) total FROM reservations WHERE 1=1 AND state = 10 GROUP BY 1",
    []string{"%d,","%d"},
    "Get new reservations by week of year",
  )

  // Daily accepted reservations
  RegisterReport("funnel_2",
    "SELECT WEEKOFYEAR(start_date) date, COUNT(id) total FROM reservations WHERE 1=1 AND state = 30 GROUP BY 1",
    []string{"%d,","%d"},
    "Get accepted reservations by week of year",
  )

  // Daily paid reservations
  RegisterReport("funnel_3",
    "SELECT WEEKOFYEAR(start_date) date, COUNT(id) total FROM reservations WHERE 1=1 AND state = 40 GROUP BY 1",
    []string{"%d,","%d"},
    "Get paid reservations by week of year",
  )
}
