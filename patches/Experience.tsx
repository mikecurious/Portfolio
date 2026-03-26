import { motion } from "framer-motion";
import { Briefcase } from "lucide-react";

const experiences = [
  {
    company: "Zueron Consulting",
    role: "IT Consultant",
    period: "2025 – Present",
    description: [
      "Provide IT consultancy services to fintech companies, including systems development, infrastructure management, and regulatory compliance advisory",
      "Conduct independent systems audits and gap assessments for financial services clients",
      "Develop IT policies, security documentation, and compliance frameworks for client submissions to regulatory bodies",
    ],
  },
  {
    company: "Roven Capital",
    role: "Systems Architect & Administration",
    period: "2025 – Present",
    description: [
      "Designed and implemented scalable, highly available cloud architectures across AWS, Azure, and GCP, leveraging load balancers, auto-scaling groups, managed databases, and object storage",
      "Administered Linux and Windows server environments, including patch management, performance tuning, storage provisioning, and scheduled maintenance",
      "Architected and enforced network segmentation, IAM policies, and role-based access controls across cloud and on-premises environments, aligning infrastructure with zero-trust security principles",
      "Led capacity planning and cost optimization initiatives across multi-cloud environments using reserved instances, right-sizing recommendations, and tagging strategies",
    ],
  },
  {
    company: "Mamlaka Hub & Spoke",
    role: "Cloud Engineer / DevOps / System Administrator",
    period: "2023 – 2025",
    description: [
      "Performed PCI DSS compliance reviews and continuous monitoring of cardholder data environments, tracking remediation of identified gaps and supporting audit readiness",
      "Diagnosed and resolved complex network issues spanning VPN connectivity, firewall rule configurations, and secure site-to-site tunnels linking corporate offices with financial institutions",
      "Deployed and managed observability stacks (Prometheus, Grafana, ELK) to enable proactive detection of infrastructure anomalies and reduce mean time to resolution",
      "Engineered and maintained CI/CD pipelines, containerized workloads (Docker), and Infrastructure-as-Code templates (Terraform, Ansible) for reliable, automated software deployments",
      "Architected and optimized cloud infrastructure across AWS, Azure, and GCP, driving resource efficiency and cost rationalization across multi-cloud environments",
      "Administered Active Directory environments, including user account provisioning, periodic access reviews, and cleanup exercises to enforce least-privilege principles",
    ],
  },
  {
    company: "Yet Kenya Limited",
    role: "Customer Support / DevOps / System Administrator",
    period: "2022 – 2025",
    description: [
      "Provided first-line and escalated desktop support, diagnosing and resolving hardware, software, and connectivity issues for end users across the organization",
      "Troubleshot and resolved network issues including VPN connectivity, firewall rules, and secure site-to-site connections between corporate offices and financial institutions",
      "Configured and maintained cloud infrastructure across AWS, Azure, and GCP; optimized resource utilization and developed CI/CD pipelines, Docker containers, and IaC scripts (Terraform, Ansible)",
      "Conducted regular security audits, penetration tests, and vulnerability assessments; implemented remediation measures",
      "Performed external system audits on client companies to ensure IT compliance with regulatory standards",
      "Mentored junior engineers in IT best practices and conducted training sessions",
    ],
  },
  {
    company: "Animal Health & Industry Training Institute (AHITI Kabete)",
    role: "IT Support",
    period: "2020 – 2022",
    description: [
      "Provided computer troubleshooting and technical support for staff and institutional systems",
      "Performed data entry, record management, and KVB indexing with high accuracy",
      "Resolved day-to-day IT issues including printer setup, software installation, and network connectivity",
    ],
  },
  {
    company: "Comet Designers",
    role: "Customer Service & Office Support",
    period: "2018 – 2020",
    description: [
      "Provided customer service and handled client inquiries on printing and design requests",
      "Maintained organized records and documentation for office operations",
    ],
  },
];

const Experience = () => {
  return (
    <section id="experience" className="py-24">
      <div className="container mx-auto px-4">
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
        >
          <h2 className="text-3xl font-bold mb-2 text-gradient">Work Experience</h2>
          <div className="w-20 h-1 bg-gradient-to-r from-primary to-accent mb-12 rounded-full" />
        </motion.div>

        <div className="relative">
          {/* Timeline line */}
          <div className="absolute left-4 md:left-1/2 top-0 bottom-0 w-px bg-border md:-translate-x-px" />

          <div className="space-y-12">
            {experiences.map((exp, i) => (
              <motion.div
                key={exp.company}
                initial={{ opacity: 0, x: i % 2 === 0 ? -30 : 30 }}
                whileInView={{ opacity: 1, x: 0 }}
                viewport={{ once: true }}
                transition={{ delay: 0.2 }}
                className={`relative flex flex-col md:flex-row ${
                  i % 2 === 0 ? "md:flex-row" : "md:flex-row-reverse"
                } items-start md:items-center gap-8`}
              >
                {/* Timeline dot */}
                <div className="absolute left-4 md:left-1/2 w-3 h-3 bg-primary rounded-full -translate-x-1.5 mt-2 md:mt-0 glow-primary z-10" />

                <div className={`ml-12 md:ml-0 md:w-1/2 ${i % 2 === 0 ? "md:pr-12 md:text-right" : "md:pl-12"}`}>
                  <div className="bg-card border border-border rounded-lg p-6 hover:border-primary/30 transition-colors">
                    <div className="flex items-center gap-2 mb-2 text-primary">
                      <Briefcase size={16} />
                      <span className="font-mono text-xs">{exp.period}</span>
                    </div>
                    <h3 className="text-lg font-bold text-foreground mb-1">{exp.role}</h3>
                    <p className="text-sm text-accent font-mono mb-4">{exp.company}</p>
                    <ul className={`space-y-2 ${i % 2 === 0 ? "md:text-left" : ""}`}>
                      {exp.description.map((item, j) => (
                        <li key={j} className="text-sm text-muted-foreground flex items-start gap-2">
                          <span className="text-primary mt-1 text-xs shrink-0">▹</span>
                          {item}
                        </li>
                      ))}
                    </ul>
                  </div>
                </div>
              </motion.div>
            ))}
          </div>
        </div>
      </div>
    </section>
  );
};

export default Experience;
