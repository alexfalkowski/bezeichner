# frozen_string_literal: true

When('the system requests the {string} with HTTP') do |name|
  opts = {
    headers: { request_id: SecureRandom.uuid, content_type: :json, accept: :json },
    read_timeout: 10, open_timeout: 10
  }

  @response = Bezeichner.observability.send(name, opts)
end

Then('the system should respond with a healthy status with HTTP') do
  expect(@response.code).to eq(200)
  expect(JSON.parse(@response.body)).to eq('status' => 'SERVING')
end

Then('the system should respond with an unhealthy status with HTTP') do
  expect(@response.code).to eq(503)

  resp = JSON.parse(@response.body)
  expect(resp['status']).to eq('NOT_SERVING')
end

Then('the system should respond with metrics') do
  expect(@response.code).to eq(200)
  expect(@response.body).to include('go_info')
end

Then('I should see {string} as healthy') do |service|
  opts = {
    headers: { request_id: SecureRandom.uuid, content_type: :json, accept: :json },
    read_timeout: 10, open_timeout: 10
  }

  wait_for do
    @response = Bezeichner.observability.health(opts)
    @response.code
  end.to eq(200)

  expect(@response.body).to_not include(service)
end
