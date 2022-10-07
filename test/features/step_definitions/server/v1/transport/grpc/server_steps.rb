# frozen_string_literal: true

When('I request to identifiers with gRPC:') do |table|
  rows = table.rows_hash
  @request_id = SecureRandom.uuid
  metadata = { 'request-id' => @request_id, 'ua' => Bezeichner.server_config['transport']['grpc']['user_agent'] }

  request = Bezeichner::V1::GetIdentifiersRequest.new(application: rows['application'], count: rows['count'].to_i)
  @response = Bezeichner::V1.server_grpc.get_identifiers(request, { metadata: metadata })
rescue StandardError => e
  @response = e
end

Then('I should receive a successfuly identifiers from gRPC:') do |table|
  rows = table.rows_hash

  expect(@response.ids.length).to eq(rows['count'].to_i)
  expect(@response.ids.first.length).to be > 0
end

Then('I should receive a not found identifiers from gRPC') do
  expect(@response).to be_a(GRPC::NotFound)
end
