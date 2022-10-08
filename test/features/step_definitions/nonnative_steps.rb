# frozen_string_literal: true

Given('the system is having issues for the application:') do |table|
  rows = table.rows_hash
  service = Nonnative.pool.service_by_name(rows['application'])

  service.proxy.close_all
end

Then('the system should return to a healthy state for the following appliation:') do |table|
  rows = table.rows_hash
  service = Nonnative.pool.service_by_name(rows['application'])

  service.proxy.reset
end
