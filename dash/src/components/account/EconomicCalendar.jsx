import React from "react";

// Generated from https://www.investing.com/webmaster-tools/economic-calendar
const EconomicCalendar = () => (
  <div>
    <h2>US Economic Calendar</h2>
    <iframe
      title="US Economic Calendar"
      src="https://sslecal2.forexprostools.com?columns=exc_flags,exc_currency,exc_importance,exc_actual,exc_forecast,exc_previous&features=datepicker,filters&countries=5&calType=week&timeZone=8&lang=1"
      width="600"
      height="400"
      frameBorder="0"
      marginWidth="0"
      marginHeight="0"
    />
  </div>
);

export default EconomicCalendar;
