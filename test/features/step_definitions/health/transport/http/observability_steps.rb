# frozen_string_literal: true

When('the system requests the {string} with HTTP') do |name|
  @response = Nonnative.observability.send(name, Bezeichner.http_options)
end

Then('the system should respond with a healthy status with HTTP') do
  expect(@response.code).to eq(200)
  expect(@response.body.strip).to eq('SERVING')
end

Then('the system should respond with metrics') do
  expect(@response.code).to eq(200)
  expect(@response.body).to include('go_info')
end
