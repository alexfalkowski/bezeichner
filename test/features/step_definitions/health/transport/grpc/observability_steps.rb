# frozen_string_literal: true

When('the system requests the health status with gRPC') do
  request = Grpc::Health::V1::HealthCheckRequest.new(service: 'bezeichner.v1.Service')
  @response = Bezeichner.health_grpc.check(request, Bezeichner.grpc_options)
end

Then('the system should respond with a healthy status with gRPC') do
  expect(@response.status).to eq(:SERVING)
end
