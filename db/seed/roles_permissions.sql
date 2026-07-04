INSERT INTO roles (name) VALUES
    ('owner'),
    ('admin'),
    ('member'),
    ('viewer')
ON CONFLICT (name) DO NOTHING;

INSERT INTO permissions (name) VALUES
    ('users.read'),
    ('users.invite'),
    ('projects.write'),
    ('billing.read')
ON CONFLICT (name) DO NOTHING;

INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id FROM roles r, permissions p
WHERE r.name = 'owner'
ON CONFLICT DO NOTHING;

INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id FROM roles r, permissions p
WHERE r.name = 'admin'
ON CONFLICT DO NOTHING;

INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id FROM roles r, permissions p
WHERE r.name = 'member' AND p.name IN ('users.read', 'projects.write')
ON CONFLICT DO NOTHING;

INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id FROM roles r, permissions p
WHERE r.name = 'viewer' AND p.name = 'users.read'
ON CONFLICT DO NOTHING;
