# frozen_string_literal: true

Given('the system is having issues for the application {string}') do |application|
  service = Nonnative.pool.service_by_name(application)

  service.proxy.close_all
end

Then('the system should return to a healthy state for the following application {string}') do |application|
  service = Nonnative.pool.service_by_name(application)

  service.proxy.reset
end
