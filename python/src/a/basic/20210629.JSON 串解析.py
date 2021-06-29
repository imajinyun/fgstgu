# coding=utf-8

import json


def parse_json_string(s: str) -> dict:
    data = json.loads(s)
    return data


if __name__ == "__main__":
    oldstr = '{\"brief\":{\"user_name\":\"maggiezhu\",\"action_type\":\"edited\",\"ruid\":\"testuatsgmaggie\",\"activity_type\":\"Default Schedule\",\"impacted_date\":\"Mon,Fri,Sat,Sun\",\"service_name\":\"my happy lunch\",\"applied_time\":\"2:00 AM - 5:30 PM\",\"changed_section_names\":\"general,turnaround_time\",\"action_date_time\":\"Tue, 29 Jun 2021, 2:25 PM\",\"grouping_id\":\"720031624947919\"},\"old\":{\"general\":{\"service_name\":\"my happy lunch\",\"service_color\":\"#5791F0\",\"start_time\":\"2:00 AM\",\"end_time\":\"5:30 PM\",\"apply_to\":\"Mon,Tue,Wed,Thu,Fri,Sat,Sun\",\"allow_children\":\"0\"},\"capacity_settings\":{\"capacity_by_covers\":{\"enable\":\"1\",\"min_party_size\":\"1\",\"max_party_size\":\"10\",\"total_covers\":\"100\",\"max_covers_interval\":\"2\",\"max_covers_waitlisted\":\"0\"},\"capacity_by_bookings\":{\"enable\":\"0\",\"list\":[]}},\"turnaround_time\":[{\"min_party_size\":\"4\",\"max_party_size\":\"8\",\"turnaround_time\":\"1:30 am\"},{\"min_party_size\":\"9\",\"max_party_size\":\"10\",\"turnaround_time\":\"1:30 am\"}],\"online_booking_slot\":[]},\"new\":{\"general\":{\"service_name\":\"my happy lunch\",\"service_color\":\"#5791F0\",\"start_time\":\"2:00 AM\",\"end_time\":\"5:30 PM\",\"apply_to\":\"Mon,Fri,Sat,Sun\",\"allow_children\":\"0\"},\"capacity_settings\":{\"capacity_by_covers\":{\"enable\":\"1\",\"min_party_size\":\"1\",\"max_party_size\":\"10\",\"total_covers\":\"100\",\"max_covers_interval\":\"2\",\"max_covers_waitlisted\":\"0\"},\"capacity_by_bookings\":{\"enable\":\"0\",\"list\":[]}},\"turnaround_time\":[{\"min_party_size\":\"4\",\"max_party_size\":\"8\",\"turnaround_time\":\"1:30 am\"}],\"online_booking_slot\":[]}}'
    newstr = parse_json_string(oldstr)
    print(newstr['brief'])
