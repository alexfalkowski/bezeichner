# frozen_string_literal: true

When('I request to list applications with HTTP') do
  @request_id = SecureRandom.uuid
  opts = Bezeichner.http_options(
    headers: {
      request_id: @request_id, user_agent: 'Bezeichner-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json
    }
  )

  @response = Bezeichner::V1.http.list(opts)
end

Then('I should receive configured applications from HTTP:') do |table|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)
  rows = table.rows_hash

  expect(resp['meta']['requestId']).to eq(@request_id)
  expect(resp['meta']['userAgent']).to eq('Bezeichner-ruby-client/1.0 HTTP/1.0')
  expect(applications(resp['generator_applications'])).to eq(mapping(rows['generator_applications']))
  expect(resp['mapper_applications'].map { |app| app['name'] }).to eq(identifiers(rows['mapper_applications']))
  expect(resp['generator_kinds']).to eq(identifiers(rows['generator_kinds']))
  expect(resp['limits']['generate_count']).to eq(rows['generate_count'].to_i)
  expect(resp['limits']['map_ids']).to eq(rows['map_ids'].to_i)
end
