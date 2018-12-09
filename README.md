# support-bot

Provides a fair rotation algorithm to decide the next person to be on support. It maintains count of the support days of individual team members and uses that to decide based on person having fewest support days. It also takes into account not to put the same person support without a gap of atleast 2 days, irrespective of the number of support days.

Will eventually be configured on slack as a `/command`
