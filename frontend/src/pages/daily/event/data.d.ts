export type RoutineInfo = {
  id: number,
  icon: string,
  short_name: string,
  event: string,
  will_spend: number,
  spend: number,
  total_spend: number,
};

export type EventInfo = {
  date: string,
  start_time: string,
  end_time: string,
  specific_event: string,
  routine_id: number,
  spend: number,
};

export type EventAPI = {
  events: EventInfo[],
  max_end_time: string,
};

export type EditEventInfo = {
  date: string,
  start_time: string,
  end_time: string,
  specific_event: string,
  routine_id: number,
};