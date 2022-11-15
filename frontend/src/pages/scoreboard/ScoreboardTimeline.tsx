import { gql } from "@apollo/client";
import { FlexibleXYPlot, HorizontalGridLines, LineSeries, XAxis, YAxis } from "react-vis";
import "react-vis/dist/style.css";
import { ScoreboardTimelineFragment } from "../../generated";

gql`
    fragment ScoreboardTimeline on scoreboard {
        team {
            id
            name
            score_timeline {
                event_time
                score
            }
        }
    }
`;
type ScoreboardTimelineProps = {
  scoreboard?: ScoreboardTimelineFragment[];
};

export default function ScoreboardTimeline({
                                             scoreboard,
                                           }: ScoreboardTimelineProps) {
  return (
    <FlexibleXYPlot xType={"time"}>
      <XAxis />
      <YAxis />
      <HorizontalGridLines />
      {scoreboard?.map((team) => (
        <LineSeries
          key={team.team?.id}
          data={team.team?.score_timeline.map((st) => ({
            x: new Date(st.event_time).getTime(),
            y: st.score,
          }))}
        />
      ))}
    </FlexibleXYPlot>
  );
}
